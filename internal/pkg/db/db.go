package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"
)

type DB struct {
	conn PostgresConn
}

func New() (*DB, error) {
	connStr := createConnectionString(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	ctx := context.Background()

	// Create connection pool
	var conn PostgresConn
	var err error
	conn, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Test ping
	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Init() error {
	err := db.CreateTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}
	return nil
}

// CreateTables creates the high-level tables for the database
func (db *DB) CreateTables() error {
	createStatements := getCreateTableSQLCommands()
	for _, s := range createStatements {
		_, err := db.conn.Exec(context.Background(), s)
		if err != nil {
			return fmt.Errorf("error creating gamestats tables.\nstatement: %v\nerror: %w", s, err)
		}
	}
	return nil
}

func (db *DB) SaveTeam(team sportsblaze.Team) error {
	bytes, err := json.Marshal(transformTeamToDBTeam(team))
	if err != nil {
		return fmt.Errorf("error marshalling team %s: %w", team.Name, err)
	}
	return db.set(nbaTeamsTable.name, team.ID, bytes)
}

func (db *DB) set(table string, key string, value []byte) error {
	sql := fmt.Sprintf(setSql, table)
	_, err := db.conn.Exec(context.Background(), sql, key, value)
	if err != nil {
		return fmt.Errorf("error setting value in table %s for key %s: %w", table, key, err)
	}
	return nil
}

// PostgresConn is an interface that defines the methods we use from pgxpool.Pool
// Allows us to mock the connection in tests
type PostgresConn interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
}
