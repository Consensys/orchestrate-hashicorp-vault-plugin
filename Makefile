GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" | egrep -v "^\./\.go" | grep -v _test.go)

test:
	go test  ./... -cover -coverprofile=coverage.txt -covermode=atomic
build:
	go build -o build/orchestrate-hashicorp-vault-plugin -ldflags "-X main.buildDate=`date -u +\"%Y-%m-%dT%H:%M:%SZ\"` -X main.buildVersion=$(BUILD_VERSION)" -tags=prod -v
lint:
	@misspell -w $(GOFILES)
	@golangci-lint run --fix
lint-ci: ## Check linting
	@misspell -error $(GOFILES)
	@golangci-lint run