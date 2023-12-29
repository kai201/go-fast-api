SHELL := /bin/bash

PROJECT_NAME := "github.com/fast"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v /api/)

.PHONY: mod
# add missing and remove unused modules
mod:
	go mod tidy

.PHONY: fmt
# go format *.go files
fmt:
	gofmt -s -w .

.PHONY: dep
# download dependencies to the directory vendor
dep:
	go mod download


.PHONY: run
# run service
run:
	go run cmd/main.go


.PHONY: build
# build System for linux amd64 binary
build:
	@echo "building 'System', linux binary file will output to 'cmd/System'"
	@cd cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fastapi

.PHONY: clean
# clean binary file, cover.out, template file
clean:
	@rm -vrf internal/ecode/*.go.gen*
	@rm -vrf internal/routers/*.go.gen*
	@rm -vrf internal/handler/*.go.gen*
	@rm -vrf internal/service/*.go.gen*
	@echo "clean finished"

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all
