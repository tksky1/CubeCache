package cube

import (
	"cubeCache/config"
	"cubeCache/imple"
	"cubeCache/lru"
	"cubeCache/rpc"
	"fmt"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"hash/fnv"
	"sync"
	"time"
)

// Cube is a concurrency-safe cache instance with a name
type Cube struct {
	shards   []*lru.Cache
	mu       []sync.RWMutex
	lStates  []*lua.LState
	lStateMu []sync.Mutex
	*rpc.CreateCubeRequest
	keyGetterFunc sync.Map // map[string]string
	evictChan     chan string
}

func New(req *rpc.CreateCubeRequest) (*Cube, error) {
	cube := &Cube{
		shards:    make([]*lru.Cache, config.CacheShards),
		mu:        make([]sync.RWMutex, config.CacheShards),
		lStates:   make([]*lua.LState, config.CacheShards),
		evictChan: make(chan string, config.EliminateChanBufSize),
	}
	cube.CreateCubeRequest = req
	for i := range cube.shards {
		cube.shards[i] = lru.New(cube.MaxBytes/int64(config.CacheShards), cube.OnEvictedFunc, cube.evictChan)
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
		go cube.cleanExpireNode(i)
	}
	go cube.consumeEliminate()
	return cube, nil
}

func GetShardId(key string) int {
	fnv32 := fnv.New32()
	_, err := fnv32.Write([]byte(key))
	if err != nil {
		println("err fnv32 hash:", err.Error())
		return 0
	}
	return int(fnv32.Sum32()) % config.CacheShards
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

// consume elimination event from lru
func (c *Cube) consumeEliminate() {
	for {
		if c == nil || c.evictChan == nil {
			break
		}
		key := <-c.evictChan
		go c.onEliminate(key)
	}
}

func (c *Cube) onEliminate(key string) {
	shard := GetShardId(key)
	c.lStateMu[shard].Lock()
	defer c.lStateMu[shard].Unlock()
	err := imple.RegisterLuaFunc(c.lStates[shard], *c.OnEvictedFunc)
	if err != nil {
		logrus.Error("user onEliminate func err: " + err.Error() + " for key" + key)
	}
}

func (c *Cube) cleanExpireNode(shard int) {
	for {
		time.Sleep(config.ExpireCleanInterval)
		c.mu[shard].Lock()
		func() {
			defer c.mu[shard].Unlock()
			c.shards[shard].DeleteAllExpiredNode()
		}()
	}
}
