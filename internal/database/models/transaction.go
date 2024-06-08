package models

import "time"

type Transaction struct {
	ID          int       `db:"id" json:"id"`
	Amount      float64   `db:"amount" json:"amount"`
	Description string    `db:"description" json:"description"`
	Type        string    `db:"type" json:"type"`
	Date        time.Time `db:"date" json:"date"`
	UserId      int       `db:"user_id" json:"user_id"`
	CategoryId  int       `db:"category_id" json:"category_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
