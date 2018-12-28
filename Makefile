#!/bin/bash

export NOW=$(shell date --rfc-3339=ns)
export PKGS=$(shell go list ./... | grep -v vendor/)
export TEST_OPTS=-cover -race

PACKAGE = github.com/reyhanfahlevi/soap-absence
COMMIT_HASH = $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE = $(shell date +%FT%T%z)
ldflags = -X $(PACKAGE)/cmd.CommitHash=$(COMMIT_HASH) -X $(PACKAGE)/cmd.BuildDate=$(BUILD_DATE) -s -w

build-api:
	@echo "${NOW} == BUILDING API"
	@CGO_ENABLED=0 go build -ldflags '$(ldflags)' -o soap-absence-api cmd/api/main.go

build-cron:
	@echo "${NOW} == BUILDING CRON"
	@CGO_ENABLED=0 go build -ldflags '$(ldflags)' -o soap-absence-cron cmd/cron/main.go

update:
	@echo "${NOW} == UPDATING..."
	@dep ensure -v

run-api: build-api
	@echo "${NOW} == RUNNING CRON..."
	@./soap-absence-api

run-cron: build-cron
	@echo "${NOW} == RUNNING CRON..."
	@./soap-absence-cron
