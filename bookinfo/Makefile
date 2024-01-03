# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make              - default to 'build' target
#   make lint         - code analysis
#   make test         - run unit test (or plus integration test)
#   make build        - alias to build-local target
#   make build-local  - build local binary targets
#   make build-linux  - build linux binary targets
#   make container    - build containers
#   $ docker login registry -u username -p xxxxx
#   make push         - push containers
#   make clean        - clean up targets
#
# Not included but recommended targets:
#   make e2e-test
#
# The makefile is also responsible to populate project version information.
#

#
# Tweak the variables based on your project.
#

# This repo's root import path (under GOPATH).
ROOT := github.com/cloudwego/biz-demo/bookinfo

# Module name.
NAME := bookinfo

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]/$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

# Container registries.
REGISTRY ?= cr-cn-beijing.volces.com/coa-samples

# Container registry for base images.
BASE_REGISTRY ?= cr-cn-beijing.volces.com/coa-samples

# Helm chart repo
CHART_REPO ?= charts

#
# These variables should not need tweaking.
#

# It's necessary to set this because some environments don't link sh -> bash.
export SHELL := /bin/bash

# It's necessary to set the errexit flags for the bash shell.
export SHELLOPTS := errexit

# Project main package location.
CMD_DIR := ./cmd

# Project output directory.
OUTPUT_DIR := ./bin

# Build directory.
BUILD_DIR := ./build

IMAGE_NAME := $(IMAGE_PREFIX)$(NAME)$(IMAGE_SUFFIX)

# Current version of the project.
GOCOMMON     := $(shell if [ ! -f go.mod ]; then echo $(ROOT)/vendor/; fi)github.com/cloudwego/biz-demo/bookinfo
#VERSION      ?= $(shell git describe --tags --always --dirty)
BRANCH       ?= $(shell git branch | grep \* | cut -d ' ' -f2)
GITCOMMIT    ?= $(shell git rev-parse HEAD)
GITTREESTATE ?= $(if $(shell git status --porcelain),dirty,clean)
BUILDDATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION      ?= latest
appVersion   ?= $(VERSION)

# Available cpus for compiling, please refer to https://github.com/caicloud/engineering/issues/8186#issuecomment-518656946 for more information.
CPUS ?= $(shell /bin/bash hack/read_cpus_available.sh)

# Track code version with Docker Label.
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git describe --tags --always --dirty)"

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
HELM_LINT := /usr/local/bin/helm
NIRVANA := $(OUTPUT_DIR)/nirvana

# Default golang flags used in build and test
# -mod=vendor: force go to use the vendor files instead of using the `$GOPATH/pkg/mod`
# -p: the number of programs that can be run in parallel
# -count: run each test and benchmark 1 times. Set this flag to disable test cache
export GOFLAGS ?= -count=1

#
# Define all targets. At least the following commands are required:
#

# All targets.
.PHONY: lint test build container push

build: build-local

# more info about `GOGC` env: https://github.com/golangci/golangci-lint#memory-usage-of-golangci-lint
lint: $(GOLANGCI_LINT) $(HELM_LINT)
	@$(GOLANGCI_LINT) run
	@bash hack/helm-lint.sh

$(GOLANGCI_LINT):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.23.6

$(HELM_LINT):
	curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | sudo bash

test:
	@go test -race -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'

build-local:
	@go build -v -o $(OUTPUT_DIR)/$(NAME)                                  \
	  -ldflags "-s -w -X $(GOCOMMON)/version.module=$(NAME)                \
	    -X $(GOCOMMON)/version.version=$(VERSION)                          \
	    -X $(GOCOMMON)/version.branch=$(BRANCH)                            \
	    -X $(GOCOMMON)/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(GOCOMMON)/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(GOCOMMON)/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);

build-linux:
	/bin/bash -c 'GOOS=linux GOARCH=amd64 GOPATH=/go GOFLAGS="$(GOFLAGS)"  \
	  go build -v -o $(OUTPUT_DIR)/$(NAME)                                 \
	    -ldflags "-s -w -X $(GOCOMMON)/version.module=$(NAME)              \
	      -X $(GOCOMMON)/version.version=$(VERSION)                        \
	      -X $(GOCOMMON)/version.branch=$(BRANCH)                          \
	      -X $(GOCOMMON)/version.gitCommit=$(GITCOMMIT)                    \
	      -X $(GOCOMMON)/version.gitTreeState=$(GITTREESTATE)              \
	      -X $(GOCOMMON)/version.buildDate=$(BUILDDATE)"                   \
		$(CMD_DIR)'

container:
	@docker build -t $(REGISTRY)/$(IMAGE_NAME):$(VERSION)                  \
	  --label $(DOCKER_LABELS)                                             \
	  -f $(BUILD_DIR)/Dockerfile .;

push: container
	@docker push $(REGISTRY)/$(IMAGE_NAME):$(VERSION);

.PHONY: clean
clean:
	@-rm -vrf ${OUTPUT_DIR}

.PHONY: convert-images update-chart-version push-chart
convert-images: lint
	@find manifests -type f | sed -n 's/Chart.yaml//p' \
	| xargs -L1 helm template --set platformConfig.imageRegistry=REGISTRY_DOMAIN \
	| grep -Eo "REGISTRY_DOMAIN/.*:.*" | sed "s/^REGISTRY_DOMAIN\///g" \
	| tr -d "[:blank:],']{}\`\"" | sort -n | uniq
	@find manifests -type f -name "images*.list" | xargs cat | sort -n | uniq

# Update all addons Chart.yaml version and appVersion
update-chart-version:
	@find manifests -maxdepth 2 -type f -name 'Chart.yaml' \
	| xargs -L1 sed -i- "s|^version: .*|version: $(VERSION)|;s|^appVersion: .*|appVersion: $(appVersion)|"
	@find manifests -maxdepth 2 -type f -name 'Chart.yaml-' | xargs rm -f

push-chart:
	@find manifests -maxdepth 1 -mindepth 1 -type d -exec helm push {} $(CHART_REPO) \;

.PHONY: kitex
kitex:
	@rm -rf kitex_gen/
	@kitex -module $(ROOT) ./idl/productpage.thrift
	@kitex -module $(ROOT) ./idl/reviews.thrift
	@kitex -module $(ROOT) ./idl/ratings.thrift
	@kitex -module $(ROOT) ./idl/details.thrift

run-server: build-local
	@bash hack/run_server.sh

run-measure: build-local
	@bash hack/run_measure.sh

run-side-server: kitex build-local
	@bash hack/run_side_server.sh

container-multi-arch:
	   @docker buildx build -t $(REGISTRY)/$(IMAGE_NAME):$(VERSION)           \
         --label $(DOCKER_LABELS)                                             \
         --platform=linux/arm64,linux/amd64                                   \
         -f $(BUILD_DIR)/Dockerfile.multiarch . --push;


