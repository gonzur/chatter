package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func DBContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	return ctx, cancel
}

func QueryRows[T any](fmtSql string, args ...any) ([]T, error) {
	ctx, cancel := DBContext()
	defer cancel()

	rows, err := DB.Query(ctx, fmtSql, args)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

func QueryRow[T any](fmtSql string, args ...any) (*T, error) {
	item := new(T)
	ctx, cancel := DBContext()
	defer cancel()

	err := DB.QueryRow(ctx, fmtSql, args).Scan(item)

	return item, err
}

func init() {
	if cstr, ok := os.LookupEnv("DB_CONN"); ok {
		var err error
		ctx, cancel := DBContext()
		defer cancel()
		if DB, err = pgxpool.New(ctx, cstr); err != nil {
			issue := fmt.Sprintf("Failed to form pool: %s", err.Error())
			log.Println(issue)
			os.Exit(1)
		}
	}
}
