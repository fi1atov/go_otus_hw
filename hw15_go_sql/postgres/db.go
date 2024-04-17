package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func OpenPool() (*DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		"user=postgres password=postgres dbname=test_db port=5432 search_path=test_schema sslmode=disable",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("We cannot connect to database")
	}

	log.Println("successfully connected to database")
	return &DB{db}, nil
}
