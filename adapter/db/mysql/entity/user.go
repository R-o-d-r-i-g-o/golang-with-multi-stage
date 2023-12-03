package entity

import (
	util "myTestWithMultiStage/utils"
)

type User struct {
	Id    uint   `gorm:"column:id;    primaryKey; autoIncrement"`
	Name  string `gorm:"column:name;  not null"`
	Email string `gorm:"column:email; not null"`
}

func (User) TableName() string {
	return util.TB_USER
}
