#!/usr/bin/env bash
set -eu -o pipefail

source './scripts/project.sh'

APP_IMAGE=${APP_AUTHOR}/${APP_NAME}

run docker push ${APP_IMAGE}:server-v1
run docker push ${APP_IMAGE}:server-v2
run docker push ${APP_IMAGE}:client
