package cube

import (
	"cubeCache/imple"
	"cubeCache/lru"
	"cubeCache/rpc"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"hash/fnv"
	"sync"
)

// Cube is a concurrency-safe cache instance with a name
type Cube struct {
	shards   []*lru.Cache
	mu       []sync.RWMutex
	lStates  []*lua.LState
	lStateMu []sync.Mutex
	*rpc.CreateCubeRequest
	keyGetterFunc sync.Map // map[string]string
}

func New(req *rpc.CreateCubeRequest) (*Cube, error) {
	cube := &Cube{
		shards:  make([]*lru.Cache, 32),
		mu:      make([]sync.RWMutex, 32),
		lStates: make([]*lua.LState, 32),
	}
	cube.CreateCubeRequest = req
	for i := range cube.shards {
		cube.shards[i] = lru.New(cube.MaxBytes/32, cube.OnEvictedFunc)
		l := lua.NewState()
		if req.CubeInitFunc != nil {
			err := imple.RegisterLuaFunc(l, *req.CubeInitFunc)
			if err != nil {
				return nil, fmt.Errorf("create cube fail: register lua func error: " + err.Error())
			}
		}
		if req.CubeGetterFunc != nil {
			err := imple.RegisterLuaFunc(l, *req.CubeGetterFunc)
			if err != nil {
				return nil, fmt.Errorf("create cube fail: register lua func error: " + err.Error())
			}
		}
		if req.OnEvictedFunc != nil {
			err := imple.RegisterLuaFunc(l, *req.OnEvictedFunc)
			if err != nil {
				return nil, fmt.Errorf("create cube fail: register lua func error: " + err.Error())
			}
		}
		cube.lStates[i] = l
	}
	return cube, nil
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

func (c *Cube) Set(key string, value lru.CacheValue, getterFunc *string) {
	shard := GetShardId(key)
	if getterFunc != nil {
		c.keyGetterFunc.Store(key, *getterFunc)
	}
	c.mu[shard].Lock()
	defer c.mu[shard].Unlock()
	c.shards[shard].Set(key, value)
}

func (c *Cube) Get(key string) (value lru.CacheValue, ok bool) {
	shard := GetShardId(key)
	c.mu[shard].RLock()
	func() {
		defer c.mu[shard].RUnlock()
		value, ok = c.shards[shard].Get(key)
	}()

	// Cache do not have that record. Get by user func
	if !ok {
		keyGetter, keyGetterOk := c.keyGetterFunc.Load(key)
		keyGetterStr := keyGetter.(string)
		if keyGetterOk {
			c.lStateMu[shard].Lock()
			defer c.lStateMu[shard].Unlock()
			err := imple.RegisterLuaFunc(c.lStates[shard], keyGetterStr)
			if err != nil {
				return nil, false
			}
			ret, err := imple.ExecuteGetterLua(c.lStates[shard], "getterFor"+key, key)
			if err != nil {
				return nil, false
			}
			value = ret
			println(keyGetterStr)
			ok = true
		} else if c.CubeGetterFunc != nil {
			c.lStateMu[shard].Lock()
			defer c.lStateMu[shard].Unlock()
			ret, err := imple.ExecuteGetterLua(c.lStates[shard], "getter", key)
			if err != nil {
				return nil, false
			}
			value = ret
			ok = true
		}
		if ok {
			c.mu[shard].Lock()
			defer c.mu[shard].Unlock()
			c.shards[shard].Set(key, value)
		}
	}
	return value, ok
}
