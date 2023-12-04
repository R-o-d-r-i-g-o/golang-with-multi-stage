package main

import (
	"myTestWithMultiStage/adapter/api"
	"myTestWithMultiStage/adapter/db/mongo"
	"myTestWithMultiStage/adapter/db/mysql"
	"myTestWithMultiStage/env"
	"time"
)

func init() {
	environment := env.New()

	environment.Load()
	environment.Verify()
}

func main() {
	time.Local, _ = time.LoadLocation(env.Var.AppTimestamp)

	mysql.Connect()
	mongo.Connect()

	api.New().Run()
}
