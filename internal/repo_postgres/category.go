package repo_postgres

import (
	"database/sql"
	"github.com/insomnia-dreams-official/service-catalog/pkg/catalog"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) catalog.CategoryRepo {
	return &categoryRepository{db}
}

// FindForNavigation returns categories with path level 1 and 2; example "1" or "1.1"
func (r *categoryRepository) FindForNavigation() ([]catalog.Category, error) {
	var categories []catalog.Category

	rows, err := r.db.Query(`
		SELECT articul, name, path, link, full_link
		FROM category
		WHERE path SIMILAR TO '\d(.)?\d*'
		ORDER BY path;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := catalog.Category{}
		err := rows.Scan(&c.Articul, &c.Name, &c.Path, &c.Link, &c.FullLink)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
