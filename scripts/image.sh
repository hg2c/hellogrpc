#!/usr/bin/env bash
set -eu -o pipefail

docker build -t hg2c/hellogrpc:server -f Dockerfile.server .
docker build -t hg2c/hellogrpc:server-v2 -f Dockerfile.server-v2 .
docker build -t hg2c/hellogrpc:client -f Dockerfile.client .
