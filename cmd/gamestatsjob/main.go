package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
)

func main() {
	connString := db.CreateConnectionString(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Connection Successful\n")
	err = db.CreateTables(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create tables: %v\n", err)
	}
	defer conn.Close(context.Background())
}
