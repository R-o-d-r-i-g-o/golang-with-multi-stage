package env

import (
	"errors"
	"fmt"
	util "myTestWithMultiStage/utils"
	"reflect"

	parser "github.com/caarlos0/env/v6"
	variables "github.com/joho/godotenv"
)

var Var *enviroment

type enviroment struct {
	// MySql
	MySqlDatabase string `env:"MYSQL_DATABASE"`
	MySqlUsername string `env:"MYSQL_USERNAME"`
	MySqlPassword string `env:"MYSQL_PASSWORD"`
	MySqlPort     string `env:"MYSQL_PORT"`
	MySqlHost     string `env:"MYSQL_HOST"`
	MySqlConfig   string `env:"MYSQL_CONFIGS"`

	// Mongo
	MongoDatabase string `env:"MONGO_DATABASE"`
	MongoUsername string `env:"MONGO_USERNAME"`
	MongoPassword string `env:"MONGO_PASSWORD"`

	// App
	AppSyncApiPort  string `env:"APP_SYNC_API_PORT"`
	AppAsyncApiPort string `env:"APP_ASYNC_API_PORT"`
	AppTimestamp    string `env:"APP_TIMESTAMP"`
}

func (env *enviroment) Load() {
	variables.Load("./../.env")
	parser.Parse(env)
}

func (env *enviroment) Verify() error {
	entity := reflect.ValueOf(env).Elem()

	for index := 0; index < entity.NumField(); index++ {
		if entity.Field(index).Interface() == util.EMPTY_STING {
			fieldName := entity.Type().Field(index).Name

			return errors.New(fmt.Sprintf("Couldn't get environment variable for %s", fieldName))
		}
	}

	return nil
}
