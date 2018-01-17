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

golang::build() {
    local APP_NAME=${1:-$APP_NAME}
    local APP_PACKAGE=${2:-$APP_PACKAGE}

    local OUTPUT=./build

    for PLATFORM in ${APP_PLATFORMS}; do
        local GOOS=${PLATFORM%/*}
        local GOARCH=${PLATFORM#*/}

        local TARGET=${OUTPUT}/${APP_NAME}-${GOOS}-${GOARCH}
        run CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -o ${TARGET} -ldflags \"${LDFLAGS}\" ${APP_PACKAGE}
    done
}
