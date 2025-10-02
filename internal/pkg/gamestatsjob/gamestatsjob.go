package gamestatsjob

import (
	"fmt"
	"log"

	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func Run(storage Storage, statsClient StatsClient) error {
	// fetch daily box scores
	date := "2024-10-26" // TODO: make this value dynamic/range over multiple dates
	dailyBoxScores, err := statsClient.DailyBoxScores(date)
	if err != nil {
		return fmt.Errorf("sportsblaze error: %v", err)
	}
	log.Printf("got %d games for %s\n", len(dailyBoxScores.Games), date)

	err = handleGames(dailyBoxScores.Games, storage)
	if err != nil {
		return fmt.Errorf("error handling games: %v", err)
	}

	log.Println("game stats job completed successfully")
	return nil
}

type StatsClient interface {
	DailyBoxScores(date string) (*sportsblaze.DailyBoxScores, error)
}

type Storage interface {
	SaveTeam(team sportsblaze.Team) error
}
