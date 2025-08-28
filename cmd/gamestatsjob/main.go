package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func main() {
	// init client
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

	// write to file
	b, err := json.MarshalIndent(dailyBoxScores, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	boxScoresOutputFn, ok := os.LookupEnv("BOX_SCORES_OUTPUT_FILENAME")
	if ok {
		err = os.WriteFile(boxScoresOutputFn, b, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
