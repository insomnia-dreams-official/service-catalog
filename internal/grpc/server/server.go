package gserverimpl

import (
	catalog "github.com/insomnia-dreams-official/service-catalog/internal/grpc/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServerImpl struct {}

func Register(gs *grpc.Server) {
	catalog.RegisterCatalogServiceServer(gs, &GrpcServerImpl{})
	reflection.Register(gs)
}

