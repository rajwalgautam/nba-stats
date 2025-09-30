package gamestatsjob

import (
	multierror "github.com/hashicorp/go-multierror"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func handleGames(games []sportsblaze.Game, storage Storage) error {
	var me error
	for _, game := range games {
		err := processGame(game, storage)
		if err != nil {
			me = multierror.Append(me, err)
		}
	}
	return me
}

func processGame(game sportsblaze.Game, storage Storage) error {
	// handle teams
	err := storage.SaveTeam(game.Teams.Home)
	if err != nil {
		return err
	}
	err = storage.SaveTeam(game.Teams.Away)
	if err != nil {
		return err
	}

	// handle box score stats

	// handle player stats

	return nil
}
