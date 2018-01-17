#!/usr/bin/env bash
set -eu

source './scripts/project.sh'

golang::docker::run $@
