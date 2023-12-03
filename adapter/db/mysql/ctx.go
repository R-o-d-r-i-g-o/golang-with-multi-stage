package mysql

import (
	"fmt"
	"log"

	"myTestWithMultiStage/adapter/db/mysql/entity"
	"myTestWithMultiStage/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",
		env.Var.MySqlUsername,
		env.Var.MySqlPassword,
		env.Var.MySqlHost,
		env.Var.MySqlPort,
		env.Var.MySqlDatabase,
		env.Var.MySqlConfig,
	)

	client, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to MySQL database:", err)
	}

	if db, _ := client.DB(); db.Ping() != nil {
		log.Fatal("Couldn't ping MySQL database:", err)
	}

	migrate(client)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.VideoRegister{})

	DB = db
}
