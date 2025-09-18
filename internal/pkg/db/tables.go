package db

import "fmt"

var (
	// SQL statements
	createTableSQL = `create table if not exists %s
		(key %s primary key, value %s not null);`
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

var tables = []tableDef{
	{
		name:      "nba_game_box_scores",
		keyType:   "text",
		valueType: "jsonb",
	},
	{
		name:      "nba_specialty_stats",
		keyType:   "text",
		valueType: "jsonb",
	},
	{
		name:      "nba_teams",
		keyType:   "text",
		valueType: "jsonb",
	},
	{
		name:      "nba_players",
		keyType:   "uuid",
		valueType: "jsonb",
	},
}
