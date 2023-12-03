package main

import (
	"fmt"
	"myTestWithMultiStage/adapter/db/mongo"
	"myTestWithMultiStage/adapter/db/mysql"
	"myTestWithMultiStage/env"
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

	mysql.Connect()
	mongo.Connect()
}
