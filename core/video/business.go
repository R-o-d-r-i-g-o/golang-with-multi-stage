package video

type IVideoBusiness interface {
}

type videoBusiness struct {
	repository IVideoRepository
}

func NewBusiness(_repository IVideoRepository) IVideoBusiness {
	return &videoBusiness{
		repository: _repository,
	}
}
