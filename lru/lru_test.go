package lru

import (
	"testing"
)

func TestGet(t *testing.T) {
	lru := New(int64(-1), nil, nil)
	lru.Set("key1", CacheValue("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestElimination(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	max := len(k1 + k2 + v1 + v2)
	lru := New(int64(max), nil, nil)
	lru.Set(k1, CacheValue(v1))
	lru.Set(k2, CacheValue(v2))
	lru.Set(k3, CacheValue(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Eliminate key1 failed")
	}
}
