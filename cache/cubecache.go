package cache

import (
	"cubeCache/cube"
	"cubeCache/rpc"
)

// CubeCache keeps map from name to cube
type CubeCache struct {
	Cubes map[string]*cube.Cube
}

func New() *CubeCache {
	return &CubeCache{Cubes: make(map[string]*cube.Cube)}
}

func (c *CubeCache) NewCube(request *rpc.CreateCubeRequest) {
	newCube := cube.New(request)
	c.Cubes[request.CubeName] = newCube
}
