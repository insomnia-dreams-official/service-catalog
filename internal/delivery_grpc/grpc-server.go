package delivery_grpc

import (
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
	pb "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	//pb.UnimplementedCatalogServer
	navigationUcase catalog.NavigationUcase
}

func Register(gs *grpc.Server, navigationUcase catalog.NavigationUcase) { // todo: think how to create function with rest/spread
	pb.RegisterCatalogServer(gs, &grpcServer{
		navigationUcase,
	})
	reflection.Register(gs)
}
