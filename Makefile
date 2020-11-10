.DELETE_ON_ERROR:

test:
	go test  ./... -cover -coverprofile=coverage.txt -covermode=atomic
build:
	go build -o build/orchestrate-hashicorp-vault-plugin -ldflags "-X main.buildDate=`date -u +\"%Y-%m-%dT%H:%M:%SZ\"` -X main.buildVersion=$(BUILD_VERSION)" -tags=prod -v
