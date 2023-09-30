# Variables
GOCMD = go
GOBUILD = $(GOCMD) build
GOLINT = golangci-lint

# Binary File Name 
BINARYNAME = myapp

.DEFAULT_GOAL := build

lint:
	$(GOLINT) run ./...
.PHONY:lint

build: lint 
	$(GOBUILD) -o $(BINARYNAME)
.PHONY:build