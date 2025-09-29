package gamestatsjob

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/rajwalgautam/nba-stats/internal/pkg/db"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

var (
	statsdb     DBClient
	statsClient SportsblazeClient
)

func Run() error {
	err := initDb()
	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}
	log.Println("db initialized successfully")

	err = initStatsClient()
	if err != nil {
		return fmt.Errorf("error initializing stats client: %v", err)
	}
	log.Println("stats client initialized successfully")

	// fetch daily box scores
	date := "2024-10-26" // TODO: make this value dynamic/range over multiple dates
	dailyBoxScores, err := statsClient.DailyBoxScores(date)
	if err != nil {
		return fmt.Errorf("sportsblaze error: %v", err)
	}
	log.Printf("got %d games for %s\n", len(dailyBoxScores.Games), date)

	err = handleGames(dailyBoxScores.Games)
	if err != nil {
		return fmt.Errorf("error handling games: %v", err)
	}

	log.Println("game stats job completed successfully")
	return nil
}

func initDb() error {
	var err error
	statsdb, err = db.New()
	if err != nil {
		return fmt.Errorf("db connection err: %v", err)
	}
	err = statsdb.Init()
	if err != nil {
		return fmt.Errorf("db init err: %v", err)
	}
	return nil
}

func initStatsClient() error {
	apiKey, ok := os.LookupEnv("SPORTSBLAZE_API_KEY")
	if !ok {
		return errors.New("error: sportsblaze api key required")
	}
	statsClient = sportsblaze.New(sportsblaze.Options{ApiKey: apiKey})
	return nil
}

type SportsblazeClient interface {
	DailyBoxScores(date string) (*sportsblaze.DailyBoxScores, error)
}

type DBClient interface {
	Init() error
	SaveTeam(team sportsblaze.Team) error
}
