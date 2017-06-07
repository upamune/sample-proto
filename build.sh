#!/bin/bash

GOOS=windows GOARCH=amd64 go build  -ldflags="-w -s" -o bin/grpc_windows
GOOS=darwin GOARCH=amd64 go build  -ldflags="-w -s" -o bin/grpc_darwin
GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o bin/grpc_linux

