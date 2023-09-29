package http

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/cube/:subPath", handleGet)
	r.POST("/cube/:subPath", handlePost)
	return r
}
