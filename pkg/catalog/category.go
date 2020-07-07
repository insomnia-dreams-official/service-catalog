package catalog

// Structure for category; category is a feature, which every product has.
type Category struct {
	Articul  string
	Name     string
	Path     string
	Link     string
	FullLink string
}

// CategoryUcase interface describe actions that can be done with category
type CategoryUcase interface {
	// Find categories of path's level 1
	GetRootcategories() ([]*Category, error)
	// Find direct childs of given link's category
	GetCategoryChilds(link string) ([]*Category, error)
}

// CategoryRepo interface describe actions that can be done against persistent store (postgres)
type CategoryRepo interface {
	// Find categories of path's levels 1 and 2
	FindForNavigation() ([]*Category, error)
	// Find categories of path's level 1
	GetRootcategories() ([]*Category, error)
	// Find direct childs of given link's category
	GetCategoryChilds(link string) ([]*Category, error)
}

