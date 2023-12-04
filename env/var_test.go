package env

import (
	"os"
	"testing"
)

func TestEnvironment_New(t *testing.T) {
	entity := New()

	if entity == nil {
		t.Errorf("Expected variables enity to has values, but got 'null'")
	}
}

func TestEnvironment_Load(t *testing.T) {
	err := os.Setenv("MYSQL_DATABASE", "testdb")
	if err != nil {
		t.Fatal(err)
	}

	env := &environment{}
	env.Load()

	if env.MySqlDatabase != "testdb" {
		t.Errorf("Expected MySQLDatabase to be 'testdb', got %s", env.MySqlDatabase)
	}
}
