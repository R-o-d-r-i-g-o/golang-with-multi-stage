package video

import "gorm.io/gorm"

type IVideoRepository interface {
}

type videoRepository struct {
	db *gorm.DB
}

func NewRepository(_db *gorm.DB) IVideoRepository {
	return &videoRepository{db: _db}
}

func (v *videoRepository) CreateVideo()

func (v *videoRepository) GetVideoInfo()
