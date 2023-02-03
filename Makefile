.PHONY: build
build:
		go build -v ./cmd/go-sql

.DEFUALT_GOAL := build