SHELL := /bin/bash

#!make

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
RELEASE_DIR := $(SRC_DIR)/release
CONFIG_LOCATION := $(SRC_DIR)
GO111MODULE := on
VERSION := 1.1.0

$(@info $(shell mkdir -p $(OUT_DIR) $(BIN_DIR) $(RELEASE_DIR))

build:
	go build -o $(BIN_DIR)/server $(SRC_DIR)/cmd/server/main.go

run:
	go run $(SRC_DIR)/cmd/server/main.go

compile:
	GOOS=linux GOARCH=386 go build -o $(BIN_DIR)/linux_i386/siwi $(SRC_DIR)/cmd/server/main.go
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux_x86_64/siwi $(SRC_DIR)/cmd/server/main.go
	GOOS=darwin GOARCH=386 go build -o $(BIN_DIR)/darwin_i386/siwi $(SRC_DIR)/cmd/server/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/darwin_x86_64/siwi $(SRC_DIR)/cmd/server/main.go
	GOOS=windows GOARCH=386 go build -o $(BIN_DIR)/windows_i386/siwi.exe $(SRC_DIR)/cmd/server/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/windows_x86_64/siwi.exe $(SRC_DIR)/cmd/server/main.go

release: compile
	mkdir -p $(RELEASE_DIR)
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_linux_i386.zip $(BIN_DIR)/linux_i386/siwi
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_linux_x86_64.zip $(BIN_DIR)/linux_x86_64/siwi
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_darwin_i386.zip $(BIN_DIR)/darwin_i386/siwi
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_darwin_x86_64.zip $(BIN_DIR)/darwin_x86_64/siwi
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_windows_i386.zip $(BIN_DIR)/windows_i386/siwi.exe
	zip -rj $(RELEASE_DIR)/siwi_$(VERSION)_windows_x86_64.zip $(BIN_DIR)/windows_x86_64/siwi.exe
