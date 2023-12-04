package main

import (
	"myTestWithMultiStage/adapter/api"
	"myTestWithMultiStage/adapter/db/mongo"
	"myTestWithMultiStage/adapter/db/mysql"
	"myTestWithMultiStage/env"

	_ "myTestWithMultiStage/docs"

	"time"
)

func init() {
	environment := env.New()

	environment.Load()
	environment.Verify()
}

// @title Your API Title
// @version 1.0
// @description Your API description
// @termsOfService https://example.com/terms/
// @host localhost:8081
// @basePath /api
// @schemes http https
func main() {
	time.Local, _ = time.LoadLocation(env.Var.AppTimestamp)

	mysql.Connect()
	mongo.Connect()

	api.New().Run()
}
