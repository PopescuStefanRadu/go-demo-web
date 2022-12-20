package main

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "")
	if err != nil {
		panic(err)
	}

	config, err := pgx.ParseConfig("")
	stdlib.OpenDB(*config)
}
