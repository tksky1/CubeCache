package http

import (
	"cubeCache/cache"
)

var cubeCache *cache.CubeCache

func Start() {
	r := initRouter()
	err := r.Run(":40002")
	if err != nil {
		println(err.Error())
	}
}
