package video

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type IVideoRepository interface {
	CreateVideo()
	GetVideoInfo()
}

type videoRepository struct {
	mysql *gorm.DB
	mongo *mongo.Client
}

func NewRepository(_mysql *gorm.DB, _mongo *mongo.Client) IVideoRepository {
	return &videoRepository{
		mysql: _mysql,
		mongo: _mongo,
	}
}

func (v *videoRepository) CreateVideo()

func (v *videoRepository) GetVideoInfo()
