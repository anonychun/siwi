SHELL := /bin/bash

#!make

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
RELEASE_DIR := $(SRC_DIR)/release
CONFIG_LOCATION := $(SRC_DIR)
GO111MODULE := on
VERSION := v0.2.2

.PHONY: run
run:
	go run $(SRC_DIR)/main.go

.PHONY: build
build:
	go build -ldflags="-s -w" -o $(BIN_DIR)/siwi $(SRC_DIR)/main.go

.PHONY: compile
compile:
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o $(BIN_DIR)/linux_i386/siwi $(SRC_DIR)/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/linux_x86_64/siwi $(SRC_DIR)/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/darwin_x86_64/siwi $(SRC_DIR)/main.go
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o $(BIN_DIR)/windows_i386/siwi.exe $(SRC_DIR)/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/windows_x86_64/siwi.exe $(SRC_DIR)/main.go

.PHONY: release
release: compile
	mkdir -p $(RELEASE_DIR)
	rm -f $(RELEASE_DIR)/*
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_linux_i386.zip $(BIN_DIR)/linux_i386/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_linux_x86_64.zip $(BIN_DIR)/linux_x86_64/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_darwin_x86_64.zip $(BIN_DIR)/darwin_x86_64/siwi
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_windows_i386.zip $(BIN_DIR)/windows_i386/siwi.exe
	zip -rjm $(RELEASE_DIR)/siwi_$(VERSION)_windows_x86_64.zip $(BIN_DIR)/windows_x86_64/siwi.exe
