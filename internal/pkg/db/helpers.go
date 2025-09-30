package db

import (
	"fmt"

	"github.com/rajwalgautam/nba-stats/internal/pkg/nbautils"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

// createConnectionString creates a connection string for connecting to postgres db
func createConnectionString(user, password, host, port, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
}

func transformTeamToDBTeam(sbteam sportsblaze.Team) Team {
	var abbrev string
	abbrev, ok := nbautils.LookupAbbrev(sbteam.Name)
	if !ok {
		abbrev = "XXX"
	}
	return Team{
		FullName:     sbteam.Name,
		Abbreviation: abbrev,
	}
}
