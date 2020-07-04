package catalog

// Structure for site navigation's item
type NavigationItem struct {
	ID    string
	Name  string
	Link  string
	Items []*NavigationItem
}

// NavigationUcase interface describe actions that can be done with navigation
type NavigationUcase interface {
	GetItems() ([]*NavigationItem, error)
}

