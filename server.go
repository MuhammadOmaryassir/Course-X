package main

import (
	"cloudX/repository"
	"cloudX/service"

	"./controller"
	router "./http"
)

var (
	videoRepository repository.VideoRepository = repository.NewDynamoDBRepository()
	videoService    service.VideoService       = service.NewVideoService(videoRepository)
	videoController controller.VideoController = controller.NewVideoController(videoService)
	httpRouter      router.Router              = router.NewGinRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/api/videos", videoController.GetVideos)
	httpRouter.POST("/api/videos", videoController.AddVideo)

	httpRouter.SERVE(port)
}
