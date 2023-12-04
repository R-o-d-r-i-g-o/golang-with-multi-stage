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

	env := &enviroment{}
	env.Load()

	if env.MySqlDatabase != "testdb" {
		t.Errorf("Expected MySQLDatabase to be 'testdb', got %s", env.MySqlDatabase)
	}
}

func TestEnvironment_Verify(t *testing.T) {
	tests := []struct {
		name           string
		env            *enviroment
		expectedErrMsg string
	}{
		{
			name: "AllVariablesSet",
			env: &enviroment{
				MySqlDatabase: "testdb",
				MySqlUsername: "root",
				MySqlPassword: "password",
				MySqlPort:     "3306",
				MySqlHost:     "localhost",
				MySqlConfig:   "charset=utf8mb4&parseTime=True&loc=Local",
				MongoDatabase: "testmongo",
				MongoUsername: "admin",
				MongoPassword: "admin_password",
				MongoHost:     "localhost",
				MongoPort:     "27017",
				AppApiPort:    "8080",
				AppTimestamp:  "2023-12-31T23:59:59Z",
			},
			expectedErrMsg: "",
		},
		{
			name: "MissingVariable",
			env: &enviroment{
				MySqlDatabase: "testdb",
				// Missing values intentionally
			},
			expectedErrMsg: "Couldn't get environment variable for MySqlUsername",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.env.Verify()

			if tt.expectedErrMsg == "" && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}

			if tt.expectedErrMsg != "" && (err == nil || err.Error() != tt.expectedErrMsg) {
				t.Errorf("Expected error: %s, but got: %v", tt.expectedErrMsg, err)
			}
		})
	}
}
