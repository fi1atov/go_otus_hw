package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool struct {
	*pgxpool.Pool
}

func OpenPool() (*DBPool, error) {
	ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/test_db?search_path=test_schema&sslmode=disable&pool_max_conns=20"

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("We couldn't find any correct DSN")
	}

	conn, err := pgxpool.NewWithConfig(ctx, pgCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal("We cannot connect to database")
	}

	log.Println("successfully connected to database")
	return &DBPool{conn}, nil
}
