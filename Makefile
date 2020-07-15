SHELL := /bin/bash

#!make

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
GO111MODULE := on
CONFIG_LOCATION := $(SRC_DIR)

$(@info $(shell mkdir -p $(OUT_DIR) $(BIN_DIR))

build:
	go build -o $(BIN_DIR)/server $(SRC_DIR)/cmd/server/main.go

run:
	go run $(SRC_DIR)/cmd/server/main.go

compile:
	GOOS=linux GOARCH=386 go build -o $(BIN_DIR)/siwi-lin32 $(SRC_DIR)/cmd/server/main.go
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/siwi-lin64 $(SRC_DIR)/cmd/server/main.go
	GOOS=windows GOARCH=386 go build -o $(BIN_DIR)/siwi-win32.exe $(SRC_DIR)/cmd/server/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/siwi-win64.exe $(SRC_DIR)/cmd/server/main.go
	GOOS=darwin GOARCH=386 go build -o $(BIN_DIR)/siwi-mac32 $(SRC_DIR)/cmd/server/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/siwi-mac64 $(SRC_DIR)/cmd/server/main.go
