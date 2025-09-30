package gamestatsjob

import (
	"errors"
	"testing"

	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
	"github.com/stretchr/testify/assert"
)

func TestRunGameStatsJob(t *testing.T) {
	tests := []struct {
		name        string
		storage     Storage
		statsClient StatsClient
		testFunc    func(t *testing.T, err error)
	}{
		{
			name:        "happy path, no error",
			storage:     mockStorage{},
			statsClient: mockStatsClient{dailyBoxScoreResp: &sportsblaze.DailyBoxScores{}},
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "db error",
			storage: mockStorage{saveErr: errors.New("some db error")},
			statsClient: mockStatsClient{
				dailyBoxScoreResp: &sportsblaze.DailyBoxScores{
					Games: []sportsblaze.Game{{}},
				},
			},
			testFunc: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "some db error")
			},
		},
		{
			name:    "stats client error",
			storage: mockStorage{},
			statsClient: mockStatsClient{
				dailyBoxScoreResp: &sportsblaze.DailyBoxScores{
					Games: []sportsblaze.Game{{}},
				},
				dailyBoxScoreErr: errors.New("some stats client error"),
			},
			testFunc: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "some stats client error")
			},
		},
	}
	for _, tc := range tests {
		got := Run(tc.storage, tc.statsClient)
		tc.testFunc(t, got)
	}
}
