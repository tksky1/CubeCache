package cache

import (
	"cubeCache/cube"
	"cubeCache/rpc"
	"sync"
)

// CubeCache keeps map from name to cube
type CubeCache struct {
	cubes map[string]*cube.Cube
	mu    sync.RWMutex
}

func New() *CubeCache {
	return &CubeCache{cubes: make(map[string]*cube.Cube)}
}

func (c *CubeCache) NewCube(request *rpc.CreateCubeRequest) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCube, err := cube.New(request)
	c.cubes[request.CubeName] = newCube
	return err
}

func (c *CubeCache) GetCube(name string) (cube *cube.Cube, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cube, ok = c.cubes[name]
	return cube, ok
}
