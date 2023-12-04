package video

import "github.com/gin-gonic/gin"

type IVideoController interface {
	CreateVideo(ctx *gin.Context)
	GetVideoInfo(ctx *gin.Context)
}

type videoController struct {
	business IVideoBusiness
}

func NewController(_business IVideoBusiness) IVideoController {
	return &videoController{
		business: _business,
	}
}

func (v *videoController) CreateVideo(ctx *gin.Context) {}

func (v *videoController) GetVideoInfo(ctx *gin.Context) {}
