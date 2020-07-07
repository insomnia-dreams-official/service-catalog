package delivery_grpc

import (
	"context"
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
	pb "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
)

// CategoryToGrpc converts domain structure to grpc structure
func (s *grpcServer) CategoryToGrpc(c *catalog.Category) *pb.Category {
	return &pb.Category{
		Articul:  c.Articul,
		Name:     c.Name,
		Path:     c.Path,
		Link:     c.Link,
		FullLink: c.FullLink,
	}
}

// GetRootcategories returns rootcategories
func (s *grpcServer) GetRootcategories(_ context.Context, _ *pb.GetRootcategoriesRequest) (*pb.GetRootcategoriesResponse, error) {
	// Get rootcategories from usecase
	rootcategories, err := s.categoryUcase.GetRootcategories()
	if err != nil {
		return nil, err
	}

	// Transform domain category to grpc structure
	var pRootcategories []*pb.Category
	for _, rc := range rootcategories {
		pRootcategories = append(pRootcategories, s.CategoryToGrpc(rc))
	}

	return &pb.GetRootcategoriesResponse{Rootcategories: pRootcategories}, nil
}

// GetCategoryChilds returns direct childs of category corresponding to the req.link
func (s *grpcServer) GetCategoryChilds(_ context.Context, req *pb.GetCategoryChildsRequest) (*pb.GetCategoryChildsResponse, error) {
	childCategories, err := s.categoryUcase.GetCategoryChilds(req.Link)
	if err != nil {
		return nil, err
	}

	// Transform domain category to grpc structure
	var pCategoryChilds []*pb.Category
	for _, rc := range childCategories {
		pCategoryChilds = append(pCategoryChilds, s.CategoryToGrpc(rc))
	}

	return &pb.GetCategoryChildsResponse{CategoryChilds: pCategoryChilds}, nil
}
