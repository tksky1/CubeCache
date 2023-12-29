package config

import "time"

const (
	// CacheShards shard count of a cube
	CacheShards int = 32
	// EliminateChanBufSize eliminate channel buffer
	EliminateChanBufSize int = 10
	// ExpirationTime value will be auto-removed after this timeout, as well as node-change log expire time
	ExpirationTime time.Duration = time.Hour
	// HeartbeatInterval of master & node
	HeartbeatInterval time.Duration = time.Second
	// ExpireCleanInterval interval of clean-goroutine for expired key
	ExpireCleanInterval time.Duration = time.Second
)
