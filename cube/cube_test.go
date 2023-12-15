package cube

import (
	"cubeCache/lru"
	"cubeCache/rpc"
	"github.com/google/uuid"
	"sync"
	"testing"
)

func TestCubeSetGet(t *testing.T) {
	syncMap := new(sync.Map)
	cube, _ := New(&rpc.CreateCubeRequest{
		CubeName:       "test",
		MaxBytes:       104857600,
		CubeInitFunc:   nil,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})

	for i := 1; i <= 10000; i++ {
		go func(t *testing.T) {
			entry := lru.CacheEntry{
				Key:   uuid.New().String(),
				Value: []byte(uuid.New().String()),
			}
			cube.Set(entry.Key, entry.Value, nil)
			syncMap.Store(entry.Key, entry.Value)
			cubeOut, ok := cube.Get(entry.Key)
			if !ok {
				t.Errorf("test concurrency fail 1 at %s", entry.Key)
				return
			}
			mapOut, _ := syncMap.Load(entry.Key)
			std := mapOut.(lru.CacheValue)
			if string(std) != string(cubeOut) {
				t.Errorf("test concurrency fail 2")
				return
			}
		}(t)
	}
}
