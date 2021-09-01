package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	databaseUrl string
}

func NewPostgreSQL(databaseUrl string) *PostgreSQL {
	return &PostgreSQL{
		databaseUrl: databaseUrl,
	}
}

func (db *PostgreSQL) Open() (*sqlx.DB, error) {
	return sqlx.Open("postgres", db.databaseUrl)
}
