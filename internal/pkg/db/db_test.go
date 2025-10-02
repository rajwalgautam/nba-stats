package db

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name     string
		conn     PostgresConn
		testFunc func(t *testing.T, err error)
	}{
		{
			name: "happy path, no error",
			conn: mockConn{},
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "error",
			conn: mockConn{execErr: errors.New("some error")},
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range tests {
		db := DB{conn: tc.conn}
		got := db.Init()
		tc.testFunc(t, got)
	}
}

func TestSaveTeam(t *testing.T) {
	tests := []struct {
		name     string
		conn     PostgresConn
		testFunc func(t *testing.T, err error)
	}{
		{
			name: "happy path, no error",
			conn: mockConn{},
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "error",
			conn: mockConn{execErr: errors.New("some error")},
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range tests {
		db := DB{conn: tc.conn}
		got := db.SaveTeam(sportsblaze.Team{})
		tc.testFunc(t, got)
	}
}

type mockConn struct {
	execErr error
	pingErr error
}

func (mc mockConn) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, mc.execErr
}

func (mc mockConn) Ping(ctx context.Context) error {
	return mc.pingErr
}
