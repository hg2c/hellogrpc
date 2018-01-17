#!/usr/bin/env bash
set -eu -o pipefail

source './scripts/project.sh'

hgc::docker::build
