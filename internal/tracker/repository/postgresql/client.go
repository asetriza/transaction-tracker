package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewClient(cfg DBConfig) (PostgreSQL, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to database: %s", err)
		return PostgreSQL{}, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %s", err)
		return PostgreSQL{}, err
	}

	return PostgreSQL{db: db}, nil
}
