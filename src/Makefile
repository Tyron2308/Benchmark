GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=BenchTesteur

all:  build test

build:
		$(GOBUILD) -o $(BINARY_NAME) -v 
test:
		@echo "no test yet"
		#$(GOTEST) -v ./...
deps:
		$(GOGET) gopkg.in/yaml.v2

