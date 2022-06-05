.PHONY: all build

BIN_DIR := ./bin
version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))


all: build


build:
	rm -f $(BIN_DIR)/osc
	go build -o $(BIN_DIR)/osc -v -ldflags "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/osc/main.go
