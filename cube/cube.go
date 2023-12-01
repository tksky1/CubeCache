package cube

import (
	"cubeCache/lru"
	"cubeCache/rpc"
	"hash/fnv"
	"sync"
)

// Cube is a concurrency-safe cache instance with a name
type Cube struct {
	shards []*lru.Cache
	mu     []sync.RWMutex
	*rpc.CreateCubeRequest
	keyGetterFunc map[string]string
}

func New(req *rpc.CreateCubeRequest) *Cube {
	cube := &Cube{
		shards: make([]*lru.Cache, 32),
		mu:     make([]sync.RWMutex, 32),
	}
	cube.CreateCubeRequest = req
	for i := range cube.shards {
		cube.shards[i] = lru.New(cube.MaxBytes/32, cube.OnEvictedFunc)
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
	if !ok {
		keyGetter, ok := c.keyGetterFunc[key]
		if ok {
			// TODO: execute keyGetter
		} else if c.CubeGetterFunc != nil {
			// TODO: execute CubeGetterFunc
		}
		if err == nil {
			value = valueOutsideCache
			ok = true
			c.shards[shard].Set(key, value)
		}
	}
	return value, ok
}
