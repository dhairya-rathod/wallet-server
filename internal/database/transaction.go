package database

import (
	"time"
)

type transaction struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	UserId      int       `json:"-"`
	CategoryId  int       `json:"category_id"`
	Category    string    `json:"category_name"`
}

func (s *Service) GetTransaction(id int) (transaction, error) {

	var txn transaction
	query := `SELECT tx.id, tx.amount, tx.description, tx.type, tx.date, c.id AS category_id, c.name AS category_name from transactions tx INNER JOIN categories c ON tx.category_id = c.id WHERE tx.id = $1;`

	row := s.db.QueryRow(query, id)

	err := row.Scan(&txn.ID, &txn.Amount, &txn.Description, &txn.Type, &txn.Date, &txn.CategoryId, &txn.Category)
	if err != nil {
		return transaction{}, err
	}

	return txn, nil
}

func (s *Service) GetTransactions(userId int) ([]transaction, error) {
	var txns []transaction
	query := `SELECT tx.id, tx.amount, tx.description, tx.type, tx.date, c.id AS category_id, c.name AS category_name from transactions tx INNER JOIN categories c ON tx.category_id = c.id WHERE tx.user_id = $1;`

	rows, err := s.db.Query(query, userId)
	if err != nil {
		return txns, err
	}
	defer rows.Close()

	for rows.Next() {
		var txn transaction
		if err := rows.Scan(&txn.ID, &txn.Amount, &txn.Description, &txn.Type, &txn.Date, &txn.CategoryId, &txn.Category); err != nil {
			return txns, err
		}
		txns = append(txns, txn)
	}

	return txns, nil
}
