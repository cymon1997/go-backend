#!/usr/bin/env bash

export NOW = $(shell date --rfc-3339=ns)

install:
	@echo "installing dependencies..."
	@go mod init github.com/cymon1997/go-backend
	@echo "install success!"

update:
	@echo "updating dependencies..."
	@go mod tidy
	@echo "update success!"

build:
	@echo "building main-app..."
	@go build -o mainapp ./cmd/mainapp/
	@echo "build success!"

run:
	@echo "starting app..."
	@./mainapp

mq:
	@echo "starting mq server..."
	@nsqlookupd & nsqd --lookupd-tcp-address=127.0.0.1:4160 & nsqadmin --lookupd-http-address 127.0.0.1:4161

all: install build run
quick: build run