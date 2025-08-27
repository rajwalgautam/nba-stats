package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// CreateTables creates the high-level tables for the database
func CreateTables(conn *pgx.Conn) error {
	createStatements := []string{
		`create table if not exists nba_game_box_scores(
	key text primary key,
	value jsonb not null);`,

		`create table if not exists nba_specialty_stats(
	key text primary key,
	value jsonb not null);`,

		`create table if not exists nba_teams(
	key text primary key,
	value jsonb not null);`,

		`create table if not exists nba_players(
	key uuid primary key,
	value jsonb not null);`,
	}
	for _, s := range createStatements {
		_, err := conn.Exec(context.Background(), s)
		if err != nil {
			return fmt.Errorf("error creating gamestats tables.\nstatement: %v\nerror: %w", s, err)
		}
	}
	log.Println("Successfully created gamestats tables")
	return nil
}
