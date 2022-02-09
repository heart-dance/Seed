SHELL := /bin/bash
BASEDIR = $(shell pwd)
APP = seed
BuildDIR = build

all: release_linux release_win release_mac

release_linux:
	# Build for linux
	go clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${BuildDIR}/${APP}-linux64-amd64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o ${BuildDIR}/${APP}-linux64-arm64 .

release_win:
	# Build for win
	go clean
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o ${BuildDIR}/${APP}-windows-amd64.exe .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -v -o ${BuildDIR}/${APP}-windows-arm.exe .

release_mac:
	# Build for mac
	go clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o ${BuildDIR}/${APP}-darwin-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o ${BuildDIR}/${APP}-darwin-arm64 .

clean:
	@go clean --cache
	@rm -rvf build/*

.PHONY: release_linux release_win release_mac clean