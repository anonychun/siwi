SHELL := /bin/bash

#!make

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
RELEASE_DIR := $(SRC_DIR)/release
CONFIG_LOCATION := $(SRC_DIR)
GO111MODULE := on
VERSION := v0.1.0

$(@info $(shell mkdir -p $(OUT_DIR) $(BIN_DIR) $(RELEASE_DIR))

install_requirements:
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	go mod tidy

dev: embed
	go run $(SRC_DIR)/main.go

build: embed
	go build -ldflags="-s -w" -o $(BIN_DIR)/siwi $(SRC_DIR)/main.go

embed:
	rm -f $(SRC_DIR)/internal/webui/rice-box.go
	cd $(SRC_DIR)/internal/webui && rice embed-go

compile: embed
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o $(BIN_DIR)/linux_i386/siwi $(SRC_DIR)/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/linux_x86_64/siwi $(SRC_DIR)/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/darwin_x86_64/siwi $(SRC_DIR)/main.go
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o $(BIN_DIR)/windows_i386/siwi.exe $(SRC_DIR)/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/windows_x86_64/siwi.exe $(SRC_DIR)/main.go

release: compile
	mkdir -p $(RELEASE_DIR)
	rm -f $(RELEASE_DIR)/*
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_linux_i386.zip $(BIN_DIR)/linux_i386/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_linux_x86_64.zip $(BIN_DIR)/linux_x86_64/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_darwin_x86_64.zip $(BIN_DIR)/darwin_x86_64/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_windows_i386.zip $(BIN_DIR)/windows_i386/siwi.exe
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_windows_x86_64.zip $(BIN_DIR)/windows_x86_64/siwi.exe
