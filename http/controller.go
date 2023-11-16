package http

import (
	"cubeCache/lru"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

func handleGet(ctx *gin.Context) {
	cubeName := ctx.Param("subPath")
	cube, ok := cubeCache.Cubes[cubeName]
	if !ok {
		ctx.JSON(400, gin.H{"msg": "cube " + cubeName + " not found"})
		return
	}
	key := ctx.Param("key")
	byteValue, ok := cube.Get(key)
	if !ok {
		ctx.JSON(401, gin.H{"msg": "user getter func for " + cubeName + "/" + key + " error"})
		return
	}
	bytes := byteValue.(*lru.Bytes)
	ctx.Data(200, "application/octet-stream", bytes.B)
}

func handlePost(ctx *gin.Context) {
	cubeName := ctx.Param("subPath")
	cube, ok := cubeCache.Cubes[cubeName]
	if !ok {
		ctx.JSON(400, gin.H{"msg": "cube " + cubeName + " not found"})
		return
	}
	key := ctx.Param("key")
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "read request body fail"})
		return
	}
	byteValue := &lru.Bytes{B: body}
	cube.Set(key, byteValue)
	ctx.JSON(200, gin.H{"msg": "success"})
}

func handleCreateCube(ctx *gin.Context) {
	name := ctx.PostForm("name")
	maxBytes, err := strconv.ParseInt(ctx.PostForm("max_bytes"), 10, 64)
	if name == "" || err != nil {
		ctx.JSON(400, gin.H{"msg": "illegal post-form params"})
		return
	}
	cubeCache.NewCube(name, nil, nil, maxBytes)
}
