package main

import (
	"log"

	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatalf("db connection err: %v", err)
	}
	err = db.Init()
	if err != nil {
		log.Fatalf("db init err: %v", err)
	}
	log.Println("db initialized successfully")
}
