package mysql

import (
	"myTestWithMultiStage/adapter/db/mysql/entity"
	util "myTestWithMultiStage/utils"

	"gorm.io/gorm"
)

func loadSeed(db *gorm.DB) {
	user := entity.User{
		Id:    1,
		Name:  "test-user",
		Email: "test-email@gmail.com",
	}

	db.Table(util.TB_USER).Create(&user)
}
