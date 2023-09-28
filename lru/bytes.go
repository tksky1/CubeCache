package lru

// Bytes implements CacheValue for byte array
type Bytes struct {
	B []byte
}

func (b *Bytes) Len() int {
	return len(b.B)
}

func (b *Bytes) String() string {
	return string(b.B)
}
