package video

type IVideoController interface {
}

type videoController struct {
	business IVideoBusiness
}

func NewController(_business IVideoBusiness) IVideoController {
	return &videoController{
		business: _business,
	}
}
