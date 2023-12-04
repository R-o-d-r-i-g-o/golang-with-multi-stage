package env

import (
	"fmt"
	"log"
	util "myTestWithMultiStage/utils"
	"reflect"

	parser "github.com/caarlos0/env/v6"
	variables "github.com/joho/godotenv"
)

var Var *environment

type environment struct {
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
	MongoHost     string `env:"MONGO_HOST"`
	MongoPort     string `env:"MONGO_PORT"`

	// App
	AppApiPort   string `env:"APP_API_PORT"`
	AppTimestamp string `env:"APP_TIMESTAMP"`
}

func New() *environment {
	return &environment{}
}

func (env *environment) Load() {
	variables.Load(".env")
	parser.Parse(env)

	Var = env
}

func (env *environment) Verify() {
	entity := reflect.ValueOf(env).Elem()

	for index := 0; index < entity.NumField(); index++ {
		if entity.Field(index).Interface() == util.EMPTY_STING {
			fieldName := entity.Type().Field(index).Name

			log.Fatal(fmt.Sprintf("Couldn't get environment variable for %s", fieldName))
		}
	}
}
