SHELL := /bin/bash

#!make

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
GO111MODULE := on
CONFIG_LOCATION := $(SRC_DIR)

$(@info $(shell mkdir -p $(OUT_DIR) $(BIN_DIR))

dev: build run

build:
	go build -o $(BIN_DIR)/server $(SRC_DIR)/cmd/server/main.go

run:
	$(BIN_DIR)/server

compile:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/siwi-lin64 $(SRC_DIR)/cmd/server/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/siwi-win64.exe $(SRC_DIR)/cmd/server/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/siwi-mac64 $(SRC_DIR)/cmd/server/main.go