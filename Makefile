ENV = /usr/bin/env
SHELL = /bin/sh

.SHELLFLAGS = -c # Run commands in a -c flag
.SILENT: ; # no need for @
.ONESHELL: ; # recipes execute in same shell
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: all # All targets are accessible for user
.DEFAULT: help # Running Make will run the help target

# select the compiler and flags
# this is useful for msan
CC = clang
CXX = clang++

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build sherlock
	go mod download
	go build -o sherlock
	echo 'Build done...'
	echo 'Generated executable sherlock'

dep: ## Get the dependencies
	go mod download

lint: dep ## Run linter
	golangci-lint  run -v

race: dep ## Run data race detector
	go test -race -short ./...

msan: dep ## Run memory sanitizer
	go test -msan -short ./... -count=1

test: dep ## Run tests
	go test ./... -v -count=1
