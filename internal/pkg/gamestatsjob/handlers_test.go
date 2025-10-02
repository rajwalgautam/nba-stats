package gamestatsjob

import (
	"errors"
	"testing"

	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
	"github.com/stretchr/testify/assert"
)

func TestProcessGame(t *testing.T) {
	tests := []struct {
		name     string
		game     sportsblaze.Game
		storage  mockStorage
		prep     func()
		testFunc func(t *testing.T, err error)
	}{
		{
			name:    "happy path, no error",
			game:    sportsblaze.Game{ID: "id1"},
			storage: mockStorage{},
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "some db save error",
			storage: mockStorage{saveErr: errors.New("some error")},
			game:    sportsblaze.Game{ID: "id1"},
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range tests {
		got := processGame(tc.game, tc.storage)
		tc.testFunc(t, got)
	}
}

func TestHandleGames(t *testing.T) {
	tests := []struct {
		name     string
		games    []sportsblaze.Game
		storage  mockStorage
		prep     func()
		testFunc func(t *testing.T, err error)
	}{
		{
			name:    "happy path, no error",
			games:   []sportsblaze.Game{{ID: "id1"}},
			storage: mockStorage{},
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "some db save error",
			storage: mockStorage{saveErr: errors.New("some error")},
			games:   []sportsblaze.Game{{ID: "id1"}},
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range tests {
		got := handleGames(tc.games, tc.storage)
		tc.testFunc(t, got)
	}
}
