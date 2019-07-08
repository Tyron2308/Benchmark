GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=BenchTesteur

all:  build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v 
test:
		@echo "no test yet"
		#$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o $(BINARY_NAME) 
		./$(BINARY_NAME)
deps:
		$(GOGET) gopkg.in/yaml.v2

