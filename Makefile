SHELL := /bin/bash

ROOT = $(shell pwd)

dep:
	go mod tidy

build-mac:
	$(MAKE) dep
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o merc main.go
