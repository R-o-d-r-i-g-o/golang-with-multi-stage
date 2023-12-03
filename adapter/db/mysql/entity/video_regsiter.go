package entity

import (
	util "myTestWithMultiStage/utils"
	"time"
)

type VideoRegister struct {
	Id              uint       `gorm:"column:id;               primaryKey; autoIncrement"`
	UploadedAt      *time.Time `gorm:"column:uploaded_at;      not null"`
	Description     string     `gorm:"column:description;      not null"`
	Author          User       `gorm:"column:author;           not null;   foreignKey:id"`
	DurationMinutes uint       `gorm:"column:duration_minutes; not null"`
}

func (VideoRegister) TableName() string {
	return util.TB_VIDEO_REGISTER
}
