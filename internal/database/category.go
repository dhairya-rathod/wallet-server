package database

import (
	"log"
	"wallet-server/internal/database/models"
)

func (s *Service) GetCategories() ([]models.Category, error) {
	var categories []models.Category

	query :=
		`SELECT id, name FROM categories WHERE is_active = true;`
	rows, err := s.db.Query(query)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var e models.Category
		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			log.Fatal("error in scan: ", err)
		}
		categories = append(categories, e)
	}

	return categories, nil
}
