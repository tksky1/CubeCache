package cube

import (
	"cubeCache/lru"
	"github.com/google/uuid"
	"sync"
	"testing"
)

func printEvicted(key string, value lru.CacheValue) {
	println("evicted:", key)
}

func generateValue(key string) (lru.CacheValue, error) {
	return []byte(uuid.New().String()), nil
}

func TestCubeSetGet(t *testing.T) {
	syncMap := new(sync.Map)
	cube := New("test", generateValue, printEvicted, 104857600)
	for i := 1; i <= 10000; i++ {
		go func(t *testing.T) {
			entry := lru.CacheEntry{
				Key:   uuid.New().String(),
				Value: []byte(uuid.New().String()),
			}
			cube.Set(entry.Key, entry.Value)
			syncMap.Store(entry.Key, entry.Value)
			cubeOut, ok := cube.Get(entry.Key)
			if !ok {
				t.Errorf("test concurrency fail 1 at %s", entry.Key)
				return
			}
			mapOut, _ := syncMap.Load(entry.Key)
			std := mapOut.(*lru.Bytes).B
			if string(std) != string(cubeOut) {
				t.Errorf("test concurrency fail 2")
				return
			}
		}(t)
	}
}
