#!/usr/bin/env bash
set -eu

DOCKER_WORKDIR=/go/src/$APP_PACKAGE
DOCKER_BUILD_IMAGE="hg2c/golang:alpine"

golang::docker::run() {
    run docker run --rm -ti \
        -w ${DOCKER_WORKDIR} \
        -v ${PROJECT_ROOT}:${DOCKER_WORKDIR} \
        $DOCKER_BUILD_IMAGE \
        ${@:-bash}
}
