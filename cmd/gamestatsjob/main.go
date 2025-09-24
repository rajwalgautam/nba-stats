package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func main() {
	// init db
	db, err := db.New()
	if err != nil {
		log.Fatalf("db connection err: %v", err)
	}
	err = db.Init()
	if err != nil {
		log.Fatalf("db init err: %v", err)
	}
	log.Println("db initialized successfully")

	// init stats client
	apiKey, ok := os.LookupEnv("SPORTSBLAZE_API_KEY")
	if !ok {
		log.Fatal("sportsblaze api key required")
	}
	statsClient := sportsblaze.New(sportsblaze.Options{ApiKey: apiKey})

	// fetch daily box scores
	date := "2024-10-24"
	dailyBoxScores, err := statsClient.DailyBoxScores(date)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got %d games for %s\n", len(dailyBoxScores.Games), date)

	boxScoresOutputFn, ok := os.LookupEnv("BOX_SCORES_OUTPUT_FILENAME")
	if ok {
		// write to file
		b, err := json.MarshalIndent(dailyBoxScores, "", "	")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(boxScoresOutputFn, b, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
