package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./service"
)

var (
	videoRepository repository.videoRepository = nil // implement and use a repository
	videoService    service.videoService       = service.NewvideoService(videoRepository)
	videoController controller.VideoController = controller.VideoController(videoService)
	httpRouter      router.Router              = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/api/videos", videoController.GetVideos)
	httpRouter.POST("/api/videos", videoController.AddVideo)

	httpRouter.SERVE(port)
}
