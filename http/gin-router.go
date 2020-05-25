package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	router = gin.Default()
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(c *gin.Context)) {
	router.GET(uri, f)
}

func (*ginRouter) POST(uri string, f func(c *gin.Context)) {
	router.POST(uri, f)
}
func (*ginRouter) PUT(uri string, f func(c *gin.Context)) {
	router.PUT(uri, f)
}
func (*ginRouter) DELETE(uri string, f func(c *gin.Context)) {
	router.DELETE(uri, f)
}

func (*ginRouter) SERVE(port string) {
	fmt.Printf("Gin HTTP server running on port %v", port)
	router.Run(port)
}
