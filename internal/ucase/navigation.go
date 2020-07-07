package ucase

import (
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
	"strings"
)

type navigationUcase struct {
	categoryRepo catalog.CategoryRepo
}

func NewNavigationUcase(categoryRepo catalog.CategoryRepo) *navigationUcase {
	return &navigationUcase{
		categoryRepo,
	}
}

// GetItems return navigation items corresponding to:
// - manually hard-coded pages
// - category pages
func (u *navigationUcase) GetItems() ([]*catalog.NavigationItem, error) {
	var items []*catalog.NavigationItem
	// Create navigation items corresponding to hard-coded pages
	items = append(items,
		&catalog.NavigationItem{ID: "1", Name: "новинки", Link: "new"},
		&catalog.NavigationItem{ID: "2", Name: "распродажа", Link: "sale"},
	)
	// Create navigation items corresponding to category pages
	categories, err := u.categoryRepo.FindForNavigation()
	if err != nil {
		return nil, err
	}

	// Iterate throw categories and create NavigationItem
	// * don't forget about nested items
	for _, c := range categories {
		if !strings.Contains(c.Path, ".") { // we need to iterate throw rootcategories; example: "1", "2"...
			item := &catalog.NavigationItem{ID: c.Articul, Name: c.Name, Link: c.FullLink}
			// Insert nested items if exists (categories of path level 2)
			for _, sc := range categories {
				if sc.Path != c.Path && strings.HasPrefix(sc.Path, c.Path) {
					nestedItem := &catalog.NavigationItem{ID: sc.Articul, Name: sc.Name, Link: sc.FullLink}
					item.Items = append(item.Items, nestedItem)
				}
			}
			items = append(items, item)
		}
	}

	return items, nil
}
