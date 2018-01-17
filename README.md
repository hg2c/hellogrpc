# hellogrpc [![Build Status](https://travis-ci.org/hg2c/hellogrpc.svg?branch=master)](https://travis-ci.org/hg2c/hellogrpc)

Hello gRPC

## development

make dev

## Compiling your protocol buffers

protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
