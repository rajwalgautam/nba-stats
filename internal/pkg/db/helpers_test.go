package db

import (
	"testing"
)

func TestCreateConnectionString(t *testing.T) {
	expectedString := "postgres://user:password@host:port/dbname"

	got := CreateConnectionString("user", "password", "host", "port", "dbname")
	want := expectedString
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
