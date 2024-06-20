package database

import (
	"fmt"
	"strings"
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

func (s *Service) PatchTransaction(id int, data types.TransactionPatchReq) error {
	query := "UPDATE transactions SET "
	params := []interface{}{}
	updates := []string{}
	paramIndex := 1

	if data.Amount != nil {
		updates = append(updates, fmt.Sprintf("amount = $%d", paramIndex))
		params = append(params, *data.Amount)
		paramIndex++
	}
	if data.Description != nil {
		updates = append(updates, fmt.Sprintf("description = $%d", paramIndex))
		params = append(params, *data.Description)
		paramIndex++
	}
	if data.Type != nil {
		updates = append(updates, fmt.Sprintf("type = $%d", paramIndex))
		params = append(params, *data.Type)
		paramIndex++
	}
	if data.Date != nil {
		updates = append(updates, fmt.Sprintf("date = $%d", paramIndex))
		params = append(params, *data.Date)
		paramIndex++
	}
	if data.CategoryId != nil {
		updates = append(updates, fmt.Sprintf("category_id = $%d", paramIndex))
		params = append(params, *data.CategoryId)
		paramIndex++
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query += strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", paramIndex)
	params = append(params, id)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	return err
}

func (s *Service) DeleteTransaction(id int) error {
	query := "DELETE FROM transactions WHERE id = $1;"
	_, err := s.db.Exec(query, id)
	return err
}
