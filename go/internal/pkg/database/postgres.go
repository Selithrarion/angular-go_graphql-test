package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var DB *pgx.Conn

func InitDB() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_USER"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	DB = conn
}

func Migrate() {}
