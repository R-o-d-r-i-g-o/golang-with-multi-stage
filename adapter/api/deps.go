package api

import (
	"myTestWithMultiStage/adapter/db/mongo"
	"myTestWithMultiStage/adapter/db/mysql"
	"myTestWithMultiStage/core/video"
)

func videoInfo() video.IVideoController {
	repository := video.NewRepository(mysql.DB, mongo.DB)
	business := video.NewBusiness(repository)
	controller := video.NewController(business)

	return controller
}
