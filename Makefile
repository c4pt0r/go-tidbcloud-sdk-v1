update:
	curl -sSL https://download.pingcap.org/tidbcloud-oas.json > tidbcloud-oas.json
	swagger generate client -f ./tidbcloud-oas.json -A go-tidbcloud
