#!/usr/bin/fish
protoc --go_out=plugins=grpc:internal/grpc/protobufs internal/grpc/protobufs/catalog.proto