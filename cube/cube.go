package cube

import (
	"cubeCache/lru"
	"hash/fnv"
	"sync"
)

// Cube is a concurrency-safe cache instance with a name
type Cube struct {
	shards []*lru.Cache
	mu     []sync.RWMutex
	name   string
	// getterFunc is the custom func to call to get the target value
	getterFunc func(key string) (value lru.CacheValue, err error)
}

func New(name string, getterFunc func(key string) (value lru.CacheValue, err error),
	OnEvicted func(key string, value lru.CacheValue), maxBytes int64) *Cube {
	cube := &Cube{
		shards:     make([]*lru.Cache, 32),
		mu:         make([]sync.RWMutex, 32),
		name:       name,
		getterFunc: getterFunc,
	}
	for i := range cube.shards {
		cube.shards[i] = lru.New(maxBytes/32, OnEvicted)
	}
	return cube
}

func GetShardId(key string) int {
	fnv32 := fnv.New32()
	_, err := fnv32.Write([]byte(key))
	if err != nil {
		println("err fnv32 hash:", err.Error())
		return 0
	}
	return int(fnv32.Sum32()) % 32
}

func (c *Cube) Set(key string, value lru.CacheValue) {
	shard := GetShardId(key)
	c.mu[shard].Lock()
	defer c.mu[shard].Unlock()
	c.shards[shard].Set(key, value)
}

func (c *Cube) Get(key string) (value lru.CacheValue, ok bool) {
	shard := GetShardId(key)
	c.mu[shard].RLock()
	defer c.mu[shard].RUnlock()
	value, ok = c.shards[shard].Get(key)
	// Cache do not have that record. Get by user func
	if !ok && c.getterFunc != nil {
		valueOutsideCache, err := c.getterFunc(key)
		if err == nil {
			value = valueOutsideCache
			ok = true
			c.shards[shard].Set(key, value)
		}
	}
	return value, ok
}
