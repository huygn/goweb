package models

import "database/sql"

type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

type Message struct {
	Text string `json:"msg"`
}