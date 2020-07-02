package gserverimpl

import (
	"context"
	catalog "github.com/insomnia-dreams-official/service-catalog/internal/grpc/protobufs"
)

func (s *GrpcServerImpl) GetNavigationItems(ctx context.Context, req *catalog.GetNavigationItemsRequest) (*catalog.GetNavigationItemsResponse, error) {
	//navigation, err := s.navigationUcase.Get()
	//if err != nil {
	//	return nil, err
	//}
	//var rProducts []*pb.Product
	//for _, dp := range dProducts {
	//	rProducts = append(rProducts, s.ProductToGrpc(dp))
	//}
	//return &catalog.GetNavigationItemsResponse{Navigation: }, err
	return nil, nil
}

