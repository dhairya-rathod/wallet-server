package database

import (
	"wallet-server/internal/types"
)

func (s *Service) GetTransaction(id int) (types.TransactionRes, error) {

	var txn types.TransactionRes
	query := `SELECT tx.id, tx.amount, tx.description, tx.type, tx.date, c.id AS category_id, c.name AS category_name from transactions tx INNER JOIN categories c ON tx.category_id = c.id WHERE tx.id = $1;`

	row := s.db.QueryRow(query, id)

	err := row.Scan(&txn.ID, &txn.Amount, &txn.Description, &txn.Type, &txn.Date, &txn.CategoryId, &txn.Category)
	if err != nil {
		return txn, err
	}

	return txn, nil
}

func (s *Service) GetTransactions(userId int) ([]types.TransactionRes, error) {
	var txns []types.TransactionRes
	query := `SELECT tx.id, tx.amount, tx.description, tx.type, tx.date, c.id AS category_id, c.name AS category_name from transactions tx INNER JOIN categories c ON tx.category_id = c.id WHERE tx.user_id = $1;`

	rows, err := s.db.Query(query, userId)
	if err != nil {
		return txns, err
	}
	defer rows.Close()

	for rows.Next() {
		var txn types.TransactionRes
		if err := rows.Scan(&txn.ID, &txn.Amount, &txn.Description, &txn.Type, &txn.Date, &txn.CategoryId, &txn.Category); err != nil {
			return txns, err
		}
		txns = append(txns, txn)
	}

	return txns, nil
}

func (s *Service) AddTransaction(data types.TransactionReq, userId int) error {
	query := "INSERT INTO transactions (user_id, amount, description, type, date, category_id) VALUES($1, $2, $3, $4, $5, $6);"

	_, err := s.db.Exec(query, userId, data.Amount, data.Description, data.Type, data.Date, data.CategoryId)
	if err != nil {
		return err
	}

	return nil
}
