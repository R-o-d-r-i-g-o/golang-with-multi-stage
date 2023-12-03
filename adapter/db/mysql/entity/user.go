package entity

import (
	util "myTestWithMultiStage/utils"
	"time"
)

type User struct {
	Id    uint       `gorm:"column:id;    primaryKey; autoIncrement"`
	Name  *time.Time `gorm:"column:name;  not null"`
	Email string     `gorm:"column:email; not null"`
}

func (User) TableName() string {
	return util.TB_USER
}
