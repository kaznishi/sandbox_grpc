#!/bin/sh

protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:proto/helloworld
