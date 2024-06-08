package models

import "time"

type Token struct {
	ID        int       `db:"id" json:"id"`
	UserId    int       `db:"user_id" json:"user_id"`
	Token     string    `db:"token" json:"token"`
	IsRevoked bool      `db:"is_revoked" json:"is_revoked"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
