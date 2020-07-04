package delivery_grpc

import (
	"context"
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
	pb "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
)

// NavigationItemToGrpc converts domain structure to grpc structure
func (s *grpcServer) NavigationItemToGrpc(ni *catalog.NavigationItem) *pb.NavigationItem {
	// Create item
	pni := pb.NavigationItem{
		Id:    ni.ID,
		Name:  ni.Name,
		Link:  ni.Link,
		Items: []*pb.NavigationItem{},
	}
	// Transform nested items
	for _, si := range ni.Items {
		pni.Items = append(pni.Items, &pb.NavigationItem{
			Id:    si.ID,
			Name:  si.Name,
			Link:  si.Link,
			Items: []*pb.NavigationItem{},
		})
	}
	return &pni
}

// GetNavigationItems handler for getting items of navigation
func (s *grpcServer) GetNavigationItems(_ context.Context, _ *pb.GetNavigationItemsRequest) (*pb.GetNavigationItemsResponse, error) {
	// Get navigation items from usecase
	navigationItems, err := s.navigationUcase.GetItems()
	if err != nil {
		return nil, err
	}

	// Transform navigation items to grpc structure
	var pNavigationItems []*pb.NavigationItem
	for _, i := range navigationItems {
		pNavigationItems = append(pNavigationItems, s.NavigationItemToGrpc(i))
	}

	return &pb.GetNavigationItemsResponse{NavigationItems: pNavigationItems}, nil
}
