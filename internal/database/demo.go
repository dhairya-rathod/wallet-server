package database

import (
	"log"

	"wallet-server/internal/database/models"
)

// GetDemoData retrieves the demo data from the database.
func (s *Service) GetDemoData() ([]models.Demo, error) {
	var demoData []models.Demo
	query := `SELECT id, title FROM public.demo`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Demo
		if err := rows.Scan(&d.ID, &d.Title); err != nil {
			log.Fatal("error in scan", err)
		}
		demoData = append(demoData, d)
	}

	return demoData, nil
}
