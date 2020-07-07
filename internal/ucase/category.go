package ucase

import (
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
)

type categoryUcase struct {
	categoryRepo catalog.CategoryRepo
}

func NewCategoryUcase(categoryRepo catalog.CategoryRepo) *categoryUcase {
	return &categoryUcase{
		categoryRepo,
	}
}

// GetRootcategories returns categories of path's level 1
func (u *categoryUcase) GetRootcategories() ([]*catalog.Category, error) {
	rootcategories, err := u.categoryRepo.GetRootcategories()
	if err != nil {
		return nil, err
	}
	return rootcategories, nil
}

// GetRootcategories returns direct childs of given link's category
func (u *categoryUcase) GetCategoryChilds(link string) ([]*catalog.Category, error) {
	categoryChilds, err := u.categoryRepo.GetCategoryChilds(link)
	if err != nil {
		return nil, err
	}
	return categoryChilds, nil
}

