package http

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/cache/:subPath", handleGet)
	r.POST("/cache/:subPath", handlePost)
	r.POST("/cube", handleCreateCube)
	return r
}
