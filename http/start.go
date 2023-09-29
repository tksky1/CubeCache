package http

import "cubeCache/cache"

var cubeCache *cache.CubeCache

func Start() {
	r := initRouter()
	cubeCache = cache.NewCubeCache()
	cubeCache.NewCube("testcube", nil, nil, 10000)
	err := r.Run(":40002")
	if err != nil {
		println(err.Error())
	}
}
