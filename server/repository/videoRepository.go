package repository

import "database/sql"

type VideoRepositoy struct {
	db *sql.DB
}