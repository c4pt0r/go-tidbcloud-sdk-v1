package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	apiclient "github.com/c4pt0r/go-tidbcloud-sdk-v1/client"
	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/cluster"
	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/project"
	"github.com/fatih/color"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/icholy/digest"
	"github.com/manifoldco/promptui"
)

var (
	publicKey  string
	privateKey string
)

func init() {
	publicKey = os.Getenv("TIDBCLOUD_PUBLIC_KEY")
	privateKey = os.Getenv("TIDBCLOUD_PRIVATE_KEY")
}

func main() {
	if publicKey == "" || privateKey == "" {
		fmt.Println("Please set TIDBCLOUD_PUBLIC_KEY and TIDBCLOUD_PRIVATE_KEY")
		os.Exit(1)
	}

	httpclient := &http.Client{
		Transport: &digest.Transport{
			Username: publicKey,
			Password: privateKey,
		},
	}

	client := apiclient.New(
		httptransport.NewWithClient("api.tidbcloud.com", "/", []string{"https"}, httpclient),
		strfmt.Default)

	// Get list of artifacts
	artifacts, err := client.Cluster.ListProviderRegions(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Print the list of clusters
	payload := artifacts.GetPayload()
	fmt.Println("Artifacts:")
	for _, item := range payload.Items {
		b, _ := json.MarshalIndent(item, "", "  ")
		fmt.Println(string(b))
	}
	fmt.Println("------------")

	// Get list of projects, and print the clusters in each project
	projects, err := client.Project.ListProjects(project.NewListProjectsParams())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Print the list of projects
	prjsPayload := projects.GetPayload()

	var projectIDs []string
	var projectNames []string
	fmt.Println("Projects:")
	for _, item := range prjsPayload.Items {
		b, _ := json.MarshalIndent(item, "", "  ")
		fmt.Println(string(b))
		projectIDs = append(projectIDs, item.ID)
		projectNames = append(projectNames, item.Name)
	}

	// Get list of clusters
	for i, projectID := range projectIDs {
		clusters, err := client.Cluster.ListClustersOfProject(cluster.NewListClustersOfProjectParams().
			WithProjectID(projectID))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Print the list of clusters
		clustersPayload := clusters.GetPayload()
		fmt.Printf("Clusters of [%s]:\n", projectNames[i])
		for _, item := range clustersPayload.Items {
			b, _ := json.MarshalIndent(item, "", "  ")
			fmt.Println(string(b))
		}
	}
	fmt.Println("------------")
	fmt.Println()

	// create a developer-tier free cluster
	fmt.Println("Creating a developer-tier cluster")
	prompt := promptui.Select{
		Label: "Select a project",
		Items: projectNames,
	}
	i, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You choose [%s] ProjectID: %s\n", result, projectIDs[i])

	// choose cluster name
	clusterName := "go-tidbcloud-sdk-v1-example"
	clusterNamePrompt := promptui.Prompt{
		Label: fmt.Sprintf("Cluster name [default:%s]", clusterName),
	}
	result, err = clusterNamePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result != "" {
		clusterName = result
	}
	fmt.Printf("Cluster name: %s\n", clusterName)

	// enter root password
	clusterRootPwd := ""
	pwdPrompt := promptui.Prompt{
		Label: "Root Password",
		Validate: func(input string) error {
			if len(input) < 6 {
				return errors.New("Password must have more than 6 characters")
			}
			return nil
		},
		Mask: '*',
	}
	clusterRootPwd, err = pwdPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// construct the cluster creation request
	clusterDefBody := &cluster.CreateClusterBody{}
	err = clusterDefBody.UnmarshalBinary([]byte(fmt.Sprintf(`{
			"name": "%s",
			"cluster_type": "DEVELOPER",
			"cloud_provider": "AWS",
			"region": "us-west-2",
			"config" : {
				"root_password": "%s",
				"ip_access_list": [
					{
						"CIDR": "0.0.0.0/0",
						"description": "Allow All"
					}
				]
			}
		}`, clusterName, clusterRootPwd)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createClusterResult, err := client.Cluster.CreateCluster(cluster.NewCreateClusterParams().
		WithProjectID(projectIDs[i]).
		WithBody(*clusterDefBody))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newClusterID := *createClusterResult.GetPayload().ID
	fmt.Printf("Cluster created, ID: %s\n", newClusterID)

	exit := make(chan *cluster.GetClusterOK)
	go func() {
		var clusterResult *cluster.GetClusterOK
		var err error
		for {
			clusterResult, err = client.Cluster.GetCluster(cluster.NewGetClusterParams().
				WithClusterID(newClusterID).
				WithProjectID(projectIDs[i]))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			s := clusterResult.GetPayload().Status.ClusterStatus
			if s == "AVAILABLE" {
				break
			}
			time.Sleep(5 * time.Second)
		}
		exit <- clusterResult
	}()

	go func() {
		// show process bar
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf(" Waiting for cluster to be ready...normally it takes 5-10 minutes")
		s.Start()
	}()

	clusterResult := <-exit

	fmt.Println(color.GreenString("\nCluster is ready, details:"))
	b, _ := json.MarshalIndent(clusterResult.GetPayload(), "", "  ")
	fmt.Println(string(b))
}
