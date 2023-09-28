package lru

import "container/list"

// Cache is the inner lru-implementation of cache. Concurrency not supported.
type Cache struct {
	maxBytes  int64
	nowBytes  int64
	cache     map[string]*list.Element
	innerList *list.List
	OnEvicted func(key string, value CacheValue)
}

type CacheEntry struct {
	Key   string
	Value CacheValue
}

// CacheValue is some type with Len()int to tell how many bytes it takes
type CacheValue interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(key string, value CacheValue)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		cache:     make(map[string]*list.Element),
		innerList: list.New(),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value CacheValue, ok bool) {
	v, ok := c.cache[key]
	if ok {
		c.innerList.MoveToBack(v)
		return v.Value.(*CacheEntry).Value, true
	}
	return nil, false
}

func (c *Cache) EliminateOldNode() {
	old := c.innerList.Front()
	if old == nil {
		return
	}
	c.innerList.Remove(old)
	oldEntry := old.Value.(*CacheEntry)
	delete(c.cache, oldEntry.Key)
	c.nowBytes -= int64(len(oldEntry.Key)) + int64(oldEntry.Value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(oldEntry.Key, oldEntry.Value)
	}
}

func (c *Cache) Set(key string, value CacheValue) {
	element, ok := c.cache[key]
	if ok {
		c.innerList.MoveToBack(element)
		entry := element.Value.(*CacheEntry)
		c.nowBytes = c.nowBytes - int64(entry.Value.Len()) + int64(value.Len())
		entry.Value = value
	} else {
		element = c.innerList.PushBack(&CacheEntry{
			Key:   key,
			Value: value,
		})
		c.cache[key] = element
		c.nowBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != -1 && c.innerList.Len() > 0 && c.nowBytes > c.maxBytes {
		c.EliminateOldNode()
	}
}

func (c *Cache) Len() int {
	return c.innerList.Len()
}
