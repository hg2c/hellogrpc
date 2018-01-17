#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

CWD=$( cd "$( dirname "${BASH_SOURCE}" )" && pwd -P )

# project dir. sample: /go/src/github.com/hg2c/swain-go
PROJECT_ROOT=$( cd $CWD && cd .. && pwd)

# project configuration file
CONFIG_FILE=${PROJECT_ROOT}/.project
if [ -s "${CONFIG_FILE}" ]; then source ${CONFIG_FILE}; fi

INFER_LANGUAGE=golang
# /go/src/github.com/ hg2c /swain-go
_HEAD=${PROJECT_ROOT%/*}
INFER_AUTHOR=${_HEAD##*/}
# /go/src/github.com/hg2c/ swain-go
INFER_NAME=${PROJECT_ROOT##*/}
# /go/src/ github.com/hg2c/swain-go
INFER_PACKAGE=github.com/$INFER_AUTHOR/$INFER_NAME

# parse the current git commit hash
COMMIT=`git rev-parse HEAD | cut -c 1-8`

# check if the current commit has a matching tag
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)

BRANCH=$(git rev-parse --abbrev-ref HEAD)

# use the matching tag as the version, if available
# VERSION=${TAG:-$BRANCH}
VERSION=${COMMIT}

# check for changed files (not untracked files)
if [[ -n "$(git diff --cached --shortstat 2> /dev/null | tail -n1)" \
         || -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ]]; then
    COMMIT="${COMMIT}-dirty"
fi

GIT_VERSION=$VERSION
GIT_COMMIT=$COMMIT
GIT_TAG=$TAG
GIT_BRANCH=$BRANCH

BUILD_DATE=$(date '+%G-%m-%d')
BUILD_TIME=$(date '+%H:%M:%S')

VERSION=${APP_VERSION:-$GIT_VERSION}

INFER_GOOS="${GOOS:-$(go env GOHOSTOS)}"
INFER_GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"
INFER_PLATFORM="${INFER_GOOS}/${INFER_GOARCH}"

APP_NAME=${APP_NAME:-$INFER_NAME}
APP_AUTHOR=${APP_AUTHOR:-$INFER_AUTHOR}
APP_LANGUAGE=golang
APP_PACKAGE=${APP_PACKAGE:-$INFER_PACKAGE}

APP_IMAGE=${APP_AUTHOR}/${APP_NAME}:${GIT_COMMIT}
APP_IMAGE_LATEST=${APP_AUTHOR}/${APP_NAME}:latest

APP_PLATFORMS=${APP_PLATFORMS:-$INFER_PLATFORM}

LDFLAGS="\
-X main.APP_NAME=$APP_NAME \
-X main.VERSION=$VERSION \
-X main.BUILD_HASH=$GIT_COMMIT \
-X main.BUILD_BRANCH=$GIT_BRANCH \
-X main.BUILD_DATE=$BUILD_DATE \
-X main.BUILD_TIME=$BUILD_TIME \
"

# TODO dry run
run() { echo RUN: $@ && eval $@; }
dryrun() { echo DRYRUN: $@; }

show() {
    local N=$1
    eval "echo $N: \$$N"
}

build() {
    local APP_NAME=${1:-$APP_NAME}
    local APP_PACKAGE=${2:-$APP_PACKAGE}

    local OUTPUT=./build

    for PLATFORM in ${APP_PLATFORMS}; do
        local GOOS=${PLATFORM%/*}
        local GOARCH=${PLATFORM#*/}

        local TARGET=${OUTPUT}/${APP_NAME}-${GOOS}-${GOARCH}
        run CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o ${TARGET} -ldflags \"${LDFLAGS}\" ${APP_PACKAGE}
    done
}

source "./scripts/${APP_LANGUAGE}.sh"

show PROJECT_ROOT
show APP_NAME
show APP_AUTHOR
show APP_PACKAGE
show VERSION
show APP_IMAGE
