package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func Connect(URI string) (Database, error) {
	pool, err := pgxpool.New(context.Background(), URI)
	if err != nil {
		return Database{}, fmt.Errorf("Connect: pgxpool.New: %w", err)
	}

	return Database{
		pool: pool,
	}, nil
}
