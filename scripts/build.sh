#!/usr/bin/env bash
set -eu -o pipefail

GOOS=linux GOARCH=amd64 go build -o build/greeter-server-linux-amd64 greeter_server/main.go
GOOS=linux GOARCH=amd64 go build -o build/greeter-client-linux-amd64 greeter_client/main.go
