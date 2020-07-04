#!/usr/bin/fish
protoc --go_out=plugins=grpc:pkg/protobuf pkg/protobuf/catalog.proto