package utils

import (
	"fmt"
	"wallet-server/internal/types"
)

func ValidateTransactionReq(transaction types.TransactionReq) error {
	if transaction.Amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	if transaction.Type == "" {
		return fmt.Errorf("type is required")
	}
	if transaction.Date.IsZero() {
		return fmt.Errorf("date is required")
	}
	if transaction.CategoryId == 0 {
		return fmt.Errorf("categoryId is required")
	}
	return nil
}
