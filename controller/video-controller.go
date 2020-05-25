package controller

import (
	"net/http"

	"../entity"
	"../service"
	"github.com/gin-gonic/gin"
)

type controller struct{}

var videoService service.VideoService

// VideoController video controller interface
type VideoController interface {
	GetVideos(ctx *gin.Context)
	AddVideo(ctx *gin.Context)
}

// NewVideoController constructor function
func NewVideoController(service service.VideoService) VideoController {
	videoService = service
	return &controller{}
}

func (*controller) GetVideos(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")
	videos, err := videoService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Error getting the videos", "Code": 500})
	}
	ctx.JSON(http.StatusOK, gin.H{"videos": videos})
}

func (*controller) AddVideo(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Error unmarshalling data", "Code": 500})
		return
	}
	err1 := videoService.Validate(&video)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": err1.Error(), "Code": 500})

		return
	}

	result, err2 := videoService.Create(&video)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Error saving the video", "Code": 500})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})

}
