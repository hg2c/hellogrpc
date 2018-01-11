#!/usr/bin/env bash
set -eu -o pipefail

GOOS=linux GOARCH=amd64 go build -o build/greeter-server-linux-amd64 github.com/hg2c/hellogrpc/greeter_server
GOOS=linux GOARCH=amd64 go build -o build/greeter-client-linux-amd64 github.com/hg2c/hellogrpc/greeter_client
# GOOS=darwin GOARCH=amd64 go build -o build/greeter-server-darwin-amd64 github.com/hg2c/hellogrpc/greeter_server
# GOOS=darwin GOARCH=amd64 go build -o build/greeter-client-darwin-amd64 github.com/hg2c/hellogrpc/greeter_client
