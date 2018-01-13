#!/usr/bin/env bash
set -eu

source './scripts/project.sh'

build greeter-server-v1 $APP_PACKAGE/greeter_server_v1
build greeter-server-v2 $APP_PACKAGE/greeter_server_v2

build greeter-client $APP_PACKAGE/greeter_client
