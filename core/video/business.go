package video

type IVideoBusiness interface {
	CreateVideo()
	GetVideoInfo()
}

type videoBusiness struct {
	repository IVideoRepository
}

func NewBusiness(_repository IVideoRepository) IVideoBusiness {
	return &videoBusiness{
		repository: _repository,
	}
}

func (v *videoBusiness) CreateVideo() {}

func (v *videoBusiness) GetVideoInfo() {}
