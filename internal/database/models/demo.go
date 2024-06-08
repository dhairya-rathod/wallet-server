package models

type Demo struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}
