#!/usr/bin/env bash
set -eu -o pipefail

PLATFORMS=${PLATFORMS:-"linux/amd64"}

OUTPUT=./build
SCRIPT_NAME=`basename "$0"`
FAILURES=""

build() {
    local NAME=$1
    local APPSOURCE=$2

    for PLATFORM in $PLATFORMS; do
        local GOOS=${PLATFORM%/*}
        local GOARCH=${PLATFORM#*/}

        local BIN_FILENAME=${OUTPUT}/${NAME}-${GOOS}-${GOARCH}

        CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} ${APPSOURCE}"
        echo "${CMD}"
        eval $CMD || FAILURES="${FAILURES} ${BIN_FILENAME}"
    done
}



build greeter-server github.com/hg2c/hellogrpc/greeter_server
build greeter-server-v2 github.com/hg2c/hellogrpc/greeter_server_v2
build greeter-client github.com/hg2c/hellogrpc/greeter_client



if [[ "${FAILURES}" != "" ]]; then
    echo ""
    echo "${SCRIPT_NAME} failed on: ${FAILURES}"
    exit 1
fi
