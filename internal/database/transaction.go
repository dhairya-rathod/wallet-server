package database

import (
	"log"
	"wallet-server/internal/database/models"
)

func (s *Service) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `SELECT * FROM transactions`

	rows, err := s.db.Query(query, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e models.Transaction
		if err := rows.Scan(&e); err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, e)
	}

	return transactions, nil
}
