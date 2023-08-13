SHELL := /bin/bash

ROOT = $(shell pwd)

dep:
	go mod tidy

build-mac:
	$(MAKE) dep
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o mercury main.go
