#!/usr/bin/env bash
set -eu -o pipefail

go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

cd src/github.com/hg2c/hellogrpc/helloworld
mockgen github.com/hg2c/hellogrpc/helloworld GreeterClient > ../mock_helloworld/hw_mock.go
