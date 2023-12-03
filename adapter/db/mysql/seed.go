package mysql

import (
	"fmt"
	"myTestWithMultiStage/adapter/db/mysql/entity"
	util "myTestWithMultiStage/utils"

	checker "github.com/jinzhu/gorm"

	"gorm.io/gorm"
)

func loadSeed(db *gorm.DB) error {
	user := entity.User{
		Id:    1,
		Name:  "test-user",
		Email: "test-email@gmail.com",
	}

	if ok := db.Table(util.TB_USER).Create(&user).Error; ok != nil {
		if checker.IsRecordNotFoundError(ok) {
			fmt.Print("User already exists")
			return nil
		}

		return fmt.Errorf("Fail during user criation: %v", ok)
	}

	return nil
}
