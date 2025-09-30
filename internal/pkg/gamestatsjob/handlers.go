package gamestatsjob

import (
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func handleGames(games []sportsblaze.Game, storage Storage) error {
	var wg sync.WaitGroup
	wg.Add(len(games))

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
	var err, errTeam error

	// handle teams
	err = storage.SaveTeam(game.Teams.Home)
	if err != nil {
		errTeam = multierror.Append(errTeam, err)
	}
	err = storage.SaveTeam(game.Teams.Away)
	if err != nil {
		errTeam = multierror.Append(errTeam, err)
	}

	// aggregate errors
	// TODO: add errors from other stats processing
	if errTeam != nil {
		err = multierror.Append(err, errTeam)
	}

	return err
}
