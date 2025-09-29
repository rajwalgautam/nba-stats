package gamestatsjob

import (
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

func handleGames(games []sportsblaze.Game) error {
	var wg sync.WaitGroup
	wg.Add(len(games))

	var me error
	for _, game := range games {
		// concurrently handle each game
		go func() {
			defer wg.Done()
			err := processGame(game)
			if err != nil {
				me = multierror.Append(me, err)
			}
		}()
	}
	// wait for all games to get processed
	wg.Wait()
	return me
}

func processGame(game sportsblaze.Game) error {
	var err, errTeam error
	var wg sync.WaitGroup

	// concurrently handle saving to each table with wg.Go()
	// handle teams
	wg.Go(func() {
		err := statsdb.SaveTeam(game.Teams.Home)
		if err != nil {
			errTeam = multierror.Append(errTeam, err)
		}
		err = statsdb.SaveTeam(game.Teams.Away)
		if err != nil {
			errTeam = multierror.Append(errTeam, err)
		}
	})

	// TODO: add wg.Go() for other stats processing

	wg.Wait()

	// aggregate errors
	// TODO: add errors from other stats processing
	if errTeam != nil {
		err = multierror.Append(err, errTeam)
	}

	return err
}
