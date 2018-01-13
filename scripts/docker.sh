#!/usr/bin/env bash
set -eu

source './scripts/project.sh'

echo "Make with Docker..."
docker run --rm -ti \
    -w /go/src/$APP_PACKAGE \
    -v `pwd`:/go/src/$APP_PACKAGE \
    $DOCKER_BUILD_IMAGE \
    ${@:-bash}
