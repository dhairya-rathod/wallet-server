package types

import "time"

type TransactionReq struct {
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	CategoryId  int       `json:"category_id"`
}

type TransactionRes struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	UserId      int       `json:"-"`
	CategoryId  int       `json:"category_id"`
	Category    string    `json:"category_name"`
}

type TransactionPatchReq struct {
	Amount      *float64   `json:"amount,omitempty"`
	Description *string    `json:"description,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	CategoryId  *int       `json:"category_id,omitempty"`
}
