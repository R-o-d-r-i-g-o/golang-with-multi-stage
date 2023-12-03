package main

import (
	"fmt"
	"myTestWithMultiStage/adapter/db/mongo"
	"myTestWithMultiStage/adapter/db/mysql"
	"myTestWithMultiStage/env"
	"time"
)

func init() {
	environment := env.Init()

	environment.Load()
	environment.Verify()
}

func main() {
	defer func() {
		fmt.Println("Connection successfully established")
	}()

	time.Local, _ = time.LoadLocation(env.Var.AppTimestamp)

	mysql.Connect()
	mongo.Connect()
}
