.DEFAULT_GOAL := help
.PHONY:
GIT_TAG         = $(strip $(shell git describe --tags --dirty --always))

DOCKER_REPO    ?= cmattoon
DOCKER_IMAGE   ?= conntrackr
DOCKER_TAG     ?= $(GIT_TAG)

DOCKER_FULLTAG  = $(DOCKER_REPO)/$(DOCKER_IMAGE):$(DOCKER_TAG)
DOCKER_LATEST   = $(DOCKER_REPO)/$(DOCKER_IMAGE):latest

CGO_ENABLED ?= 0
GO_SOURCES   = $(shell go list ./... | grep -v vendor)

PWD       = $(shell pwd)

.PHONY: help
help: ## Display help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'

.PHONY: clean
clean: ## Clean
clean:
	@rm -rf build
	find . -name '*~' -delete

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) go build -v -o build/conntrackr main.go
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -v -o build/conntrackr-linux-amd64 main.go

.PHONY: cover
cover:
	go tool cover -html cover.out

.PHONY: test
test:
	go test -v -race $(GO_SOURCES)

.PHONY: dep
dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure -vendor-only

.PHONY: container
container: ## Build Docker Container (and tag latest)
	docker build -t $(DOCKER_FULLTAG) .
	docker tag $(DOCKER_FULLTAG) $(DOCKER_LATEST)

.PHONY: push
push: ## Push FULLTAG to DockerHub
	docker push $(DOCKER_FULLTAG)
