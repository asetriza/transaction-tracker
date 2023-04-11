package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

func NewClient(cfg Config) (Client, error) {
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
		return Client{}, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %s", err)
		return Client{}, err
	}

	return Client{DB: db}, nil
}
