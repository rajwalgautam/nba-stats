package db

import (
	"fmt"
)

// createConnectionString creates a connection string for connecting to postgres db
func createConnectionString(user, password, host, port, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
}
