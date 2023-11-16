package cache

import (
	"cubeCache/cube"
	"cubeCache/lru"
)

// CubeCache keeps map from name to cube
type CubeCache struct {
	Cubes map[string]*cube.Cube
}

func New() *CubeCache {
	return &CubeCache{Cubes: make(map[string]*cube.Cube)}
}

func (c *CubeCache) NewCube(name string, getterFunc func(key string) (value lru.CacheValue, err error),
	OnEvicted func(key string, value lru.CacheValue), maxBytes int64) {
	newCube := cube.New(name, getterFunc, OnEvicted, maxBytes)
	c.Cubes[name] = newCube
}
