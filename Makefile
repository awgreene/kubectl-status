.PHONY: all build
all: build

build:
	go build -o bin/kubectl-status

install: build
	install bin/kubectl-status $(shell go env GOPATH)/bin