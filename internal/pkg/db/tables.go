package db

import "fmt"

var (
	// SQL statements
	createTableSQL = `create table if not exists %s
		(key %s primary key, value %s not null);`

	setSql = `
		insert into %s (key, value) values ($1, $2)
		on conflict (key) do update set value = EXCLUDED.value;`
)

// Build the SQL commands to create the tables
func getCreateTableSQLCommands() []string {
	commands := make([]string, len(tables))
	for i, table := range tables {
		commands[i] = fmt.Sprintf(createTableSQL, table.name, table.keyType, table.valueType)
	}
	return commands
}

// DB Tables
type tableDef struct {
	name      string
	keyType   string
	valueType string
}

var (
	nbaBoxScoreTable = tableDef{
		name:      "nba_game_box_scores",
		keyType:   "text",
		valueType: "jsonb",
	}

	nbaSpecialtyStatsTable = tableDef{
		name:      "nba_specialty_stats",
		keyType:   "text",
		valueType: "jsonb",
	}

	nbaTeamsTable = tableDef{
		name:      "nba_teams",
		keyType:   "text",
		valueType: "jsonb",
	}

	nbaPlayersTable = tableDef{
		name:      "nba_players",
		keyType:   "uuid",
		valueType: "jsonb",
	}
)

var tables = []tableDef{nbaBoxScoreTable, nbaSpecialtyStatsTable, nbaTeamsTable, nbaPlayersTable}
