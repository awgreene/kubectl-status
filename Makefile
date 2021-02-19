.PHONY: all build
all: build

GO := go
MODULE := $(shell $(GO) list -m)
VERSION := $(shell git describe --always --tags HEAD)$(and $(shell git status --porcelain),+$(shell scripts/worktree-hash.sh))


build:
	$(GO) build -ldflags '-X $(MODULE)/internal/version.Version=$(VERSION)' -o ./bin/kubectl-status .

install: build
	mv ./bin/kubectl-status $(shell go env GOPATH)/bin

unit:
	$(GO) test -coverprofile=coverage.out ./...


	# test