package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
)

func main() {
	connString := db.CreateConnectionString(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	log.Println("Successfully connected to database")
	err = db.CreateTables(conn)
	if err != nil {
		log.Fatalf("Unable to create tables: %v\n", err)
	}
}
