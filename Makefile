APP_NAME = payment-gateway
VERSION ?= SNAPSHOT

GOENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.appName=$(APP_NAME)"
PACKAGES := $(shell go list ./... | grep -v /vendor/)

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the binary
	@echo "Building binary"
	$(GOENV) go build $(LDFLAGS) -a -installsuffix cgo \
	-o $(APP_NAME) ./cmd/$(APP_NAME)

.PHONY: test
test: ## Run Tests into the packages
	@echo "Running tests"
	go test -v -covermode=atomic -coverpkg=./... -coverprofile=cover.out ./...

.PHONY: integration
integration: ## Run Tests into the packages
	@echo "Running integration tests"
	go test -v -tags integration -covermode=atomic -coverpkg=./... -coverprofile=cover.out ./...

.PHONY: all
all: test build ## Run the tests, build the binary and push the docker image

.PHONY: help
help: ## Help
	@echo "Please use 'make <target>' where <target> is ..."
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: grpc_gateway
grpc_gateway: ## Grpc Gateway
	 protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
      -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:. \
      web/proto/v1/payment-gateway.proto

.PHONY: protoc
protoc: ## Protoc Initialize
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
      -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --go_out=plugins=grpc:. \
      web/proto/v1/payment-gateway.proto

.PHONY: swagger
swagger: ## Swagger
	protoc -I/usr/local/include -I. \
      -I $(GOPATH)/src \
      -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --swagger_out=logtostderr=true:. \
      web/proto/v1/payment-gateway.proto

.PHONY: generate
generate: protoc grpc_gateway swagger