package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
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

func (db *PostgreSQL) Migrate(dbcon sqlx.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrate/migrations",
	}

	n, err := migrate.Exec(dbcon.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations.", n)
	return nil
}
