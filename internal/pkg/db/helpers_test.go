package db

import (
	"testing"

	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnectionString(t *testing.T) {
	expectedString := "postgres://user:password@host:port/dbname"

	got := createConnectionString("user", "password", "host", "port", "dbname")
	want := expectedString
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTransformTeamToDBTeam(t *testing.T) {
	tests := []struct {
		name     string
		team     sportsblaze.Team
		testFunc func(t *testing.T, team Team)
	}{
		{
			name: "happy path, abbrev found",
			team: sportsblaze.Team{Name: "Dallas Mavericks"},
			testFunc: func(t *testing.T, team Team) {
				assert.Equal(t, "DAL", team.Abbreviation)
			},
		},
		{
			name: "abbrev not found",
			team: sportsblaze.Team{Name: "Some fake team"},
			testFunc: func(t *testing.T, team Team) {
				assert.Equal(t, "XXX", team.Abbreviation)
			},
		},
	}

	for _, tc := range tests {
		got := transformTeamToDBTeam(tc.team)
		tc.testFunc(t, got)
	}
}
