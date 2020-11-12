GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" | egrep -v "^\./\.go" | grep -v _test.go)

test:
	go test  ./... -cover -coverprofile=coverage.txt -covermode=atomic
build:
	go build -o build/orchestrate-hashicorp-vault-plugin -ldflags "-X main.buildDate=`date -u +\"%Y-%m-%dT%H:%M:%SZ\"` -X main.buildVersion=$(BUILD_VERSION)" -tags=prod -v
lint-tools: ## Install linting tools
	@GO111MODULE=on go get github.com/client9/misspell/cmd/misspell@v0.3.4
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.27.0
lint:
	@misspell -w $(GOFILES)
	@golangci-lint run --fix
lint-ci: ## Check linting
	@misspell -error $(GOFILES)
	@golangci-lint run