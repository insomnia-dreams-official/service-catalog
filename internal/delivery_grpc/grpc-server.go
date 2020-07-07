package delivery_grpc

import (
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
	pb "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	navigationUcase catalog.NavigationUcase
	categoryUcase   catalog.CategoryUcase
}

// todo: think how to create function with rest/spread
func Register(gs *grpc.Server, navigationUcase catalog.NavigationUcase, categoryUcase catalog.CategoryUcase) {
	pb.RegisterCatalogServer(gs, &grpcServer{
		navigationUcase,
		categoryUcase,
	})
	reflection.Register(gs)
}
