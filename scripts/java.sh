#!/usr/bin/env bash
set -eu

DOCKER_WORKDIR=/app
DOCKER_BUILD_IMAGE=anapsix/alpine-java:8_jdk

java::docker::run() {
    run docker run --rm -ti \
        -w ${DOCKER_WORKDIR} \
        -v ${PROJECT_ROOT}:${DOCKER_WORKDIR} \
        -v ~/.m2/:/root/.m2/ \
        -v ~/.gradle/:/root/.gradle/ \
        $DOCKER_BUILD_IMAGE \
        ${@:-bash}
}
