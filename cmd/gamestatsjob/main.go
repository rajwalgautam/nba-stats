package main

import (
	"fmt"
	"log"

	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
	"github.com/rajwalgautam/nba-stats/internal/pkg/gamestatsjob"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func main() {
	db, err := initDb()
	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}
	log.Println("db initialized successfully")

	statsClient, err := initStatsClient()
	if err != nil {
		log.Fatalf("error initializing stats client: %v", err)
	}
	log.Println("stats client initialized successfully")

	err = gamestatsjob.Run(db, statsClient)
	if err != nil {
		log.Fatalf("error running gamestatsjob: %v", err)
	}
}

func initDb() (*db.DB, error) {
	statsdb, err := db.New()
	if err != nil {
		return nil, fmt.Errorf("db connection err: %v", err)
	}
	err = statsdb.Init()
	if err != nil {
		return nil, fmt.Errorf("db init err: %v", err)
	}
	return statsdb, nil
}

func initStatsClient() (*sportsblaze.Client, error) {
	// apiKey, ok := os.LookupEnv("SPORTSBLAZE_API_KEY")
	// if !ok {
	// 	return nil, errors.New("error: sportsblaze api key required")
	// }
	apiKey := "sbf1zukz3ya0hsyvyfjb06e"
	return sportsblaze.New(sportsblaze.Options{ApiKey: apiKey}), nil
}
