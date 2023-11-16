package cache

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// Mapper maintains consistency-hash wheel for all nodes
type Mapper struct {
	// number of replicated nodes on the wheel
	replicaNum int
	// nodes keeps the wheel of hashed nodeName, sorted
	nodes        []int
	mapNodesName map[int]string
	mu           sync.RWMutex
}

func NewMapper(replicaNum int) *Mapper {
	return &Mapper{mapNodesName: make(map[int]string), replicaNum: replicaNum}
}

// AddNode add a node to the wheel
func (m *Mapper) AddNode(nodes ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, v := range nodes {
		for i := 1; i <= m.replicaNum; i++ {
			hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + v)))
			m.nodes = append(m.nodes, hash)
			m.mapNodesName[hash] = v
		}
	}
	sort.Ints(m.nodes)
}

// RemoveNode remove a node from the wheel
func (m *Mapper) RemoveNode(node string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 1; i <= m.replicaNum; i++ {
		hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + node)))
		index := sort.SearchInts(m.nodes, hash)
		if index < len(m.nodes) && m.nodes[index] == hash {
			m.nodes = append(m.nodes[:index], m.nodes[index+1:]...)
			delete(m.mapNodesName, hash)
		}
	}
}

// Get the node to visit
func (m *Mapper) Get(key string) (node string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.nodes) == 0 {
		return ""
	}
	hash := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.SearchInts(m.nodes, hash)
	if idx >= len(m.nodes) {
		idx = 0
	}
	return m.mapNodesName[m.nodes[idx]]
}
