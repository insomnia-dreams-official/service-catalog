package catalog

// Structure for category; category is a feature, which every product has.
type Category struct {
	Articul  string
	Name     string
	Path     string
	Link     string
	FullLink string
}

// CategoryRepo interface describe actions that can be done against persistent store (postgres)
type CategoryRepo interface {
	FindForNavigation() ([]Category, error)
}

