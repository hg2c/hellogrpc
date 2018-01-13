#!/usr/bin/env bash
set -eu -o pipefail

source './scripts/project.sh'

APP_IMAGE=${APP_AUTHOR}/${APP_NAME}

run docker build -t ${APP_IMAGE}:server-v1 -f Dockerfile.server-v1 .
run docker build -t ${APP_IMAGE}:server-v2 -f Dockerfile.server-v2 .
run docker build -t ${APP_IMAGE}:client -f Dockerfile.client .
