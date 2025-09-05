all: build
.PHONY: all build

BINARY_NAME := start_node
BUILD_IMAGE ?= test/start_node
PUSH_IMAGE ?= test/start_node:v1.0.0

VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X 'main.BuildVersion=$(VERSION)'
GOLDFLAGS += -X 'main.BuildTime=$(BUILDTIME)'

build:
	@CGO_ENABLED=0 go build -ldflags="-s -w $(GOLDFLAGS)" -o ./bin/$(BINARY_NAME) ./cmd/main.go

container:
	docker build -t $(BUILD_IMAGE) .

push:
	docker tag $(BUILD_IMAGE) $(PUSH_IMAGE) 
	docker push $(BUILD_IMAGE)
	docker push $(PUSH_IMAGE)

coverage:
	./buildtools/coverage.sh
	./buildtools/codecov

test: 
	@cd pkg/helloworld; go test -v --race
	@cd pkg/prodcons; go test -v --race
	@cd pkg/node; go test -v --race
	@cd pkg/gossip; go test -v --race

install:
	cp ./bin/$(BINARY_NAME) /usr/local/bin
