# CubeCache 实现笔记

---

![CubeUniverse3](http://cdn.mcyou.cc/CubeUniverse3.png)

`CubeCache`是原先计划为 CubeUniverse 实现的分布式一致性缓存组件，旨在实现 CubeUniverse 系统缓存空间的线性可扩展。原计划该组件直接利用 redis 实现，但最近看到7daysgolang和godis的内存数据库实现，又产生了造轮子的冲动，决定我也来一个。

## #0x00 背景

为什么需要缓存？不管是什么样的架构，但凡涉及瓶颈与性能分析，缓存肯定是非常重要的技术设计。对于单机数据库，在面临海量请求压力，包括数据存储本身和对应的计算（如排名）等，我们是不可能每次都让请求打入数据库的，毕竟数据库要保证ACID特性，需要牺牲性能，而且在持久化IO处是非常耗时、消耗性能的。如果我们能让大量请求不需要打入数据库，在内存解决问题，就可以极大地减轻性能压力。特别是对于 CubeUniverse 这样的面向大型分布式集群的系统，除去底层Ceph本身osd的CRUD，调度信息、平台服务层的元数据、机器学习任务元数据等如果均依赖高一致性持久化维护，就都要打入集群的不同节点的不同性能的存储设备处理请求，可能造成严重的性能瓶颈。

为什么需要分布式缓存？为什么不直接给operator加个map就当缓存了？显然地，这不符合系统的设计目标，也就是要求容量能随节点扩展而扩展，就单个机子的内存和性能肯定不够应付海量数据。其次，map本身是没有并发保护的，也不能简单的拿来直接当cache。此外，还需要严谨考虑锁设计以减少性能问题。

## #0x01 技术设计

### 参考：memcached

memcached是一个开源、分布式的高性能内存缓存系统，是一个存储键值对的 HashMap，采用kv的方式在内存中保存任意数据。memcached 的架构示意图如下：

![dsadsadsadadgfdvcd](http://cdn.mcyou.cc/dsadsadsadadgfdvcd.png)

memcached 的设计逻辑是“小而美”，什么意思呢，通过分布式哈希算法，memcached 使得哈希过程在客户端即可完成，由客户端来选择具体去储存的节点，而非由集群自己协调，这样可以精致优雅地完成key在不同节点的分步。

不过，它的问题也比较明显，首先 memcached 把服务器选择完全在客户端这里实现，集群节点之间完全没有任何通信，可以说 memcached 并不是一个严格的分布式系统。这就导致客户端需要负责维护服务器列表，这就给服务器节点的可扩展性和灵活性引入了困难。此外，memchached 几乎没有对容错的处理机制，虽然对缓存来说偶然的kv丢失不是什么问题，但如果在大量扩展下出现节点宕机的情况，由客户端去通过延时判断节点宕机再去调整服务器列表可能会出现严重的时延问题，反而违背了缓存的初衷了。

### 参考：Redis 缓存模式

redis 本身定位是一个内存kv数据库，但大多数应用实际上都是拿它做缓存，而且用的非常广泛。redis 的优化很好，支持的特性也多，市场占有量也是独角兽，可以说没什么不用它的理由了。但是仍然可以参考它的设计，尝试解决一些问题。

考虑 redis 的分布式设计。redis 提供了 cluster 模式，这种模式会组织一个去中心化的集群，这个平等集群中的每个单位又配置一主多从，在这个小群体里使用 Raft 算法来维护主，而从不提供服务。由于使用非中心化架构，redis 集群是不方便去维护一些共识信息的。我们可以考虑把我们的设计改为一个中心化架构，方便我们实现一些 redis 没法或者很难去实现的特性。

考虑 redis 的缓存模式。缓存一般有三种模式。cache aside 是最常用的，业务将缓存视为一个辅助组件，同时控制缓存和DB，这里就需要业务去决定缓存和DB的访问顺序、如何访问、是否访问等。并且这样也会引入潜在的缓存与DB不一致。事实上，redis 的一致性保证是最终一致性，也就是说缓存和DB会在一段时间内产生不一致，但一段时间后会归于一致。这也就意味着业务可能在 redis 集群读到过期数据。



write through

热点key倾斜

缓存穿透：布隆过滤器

![cubecache](http://cdn.mcyou.cc/202311162353459.png)

## #0x02 实现单机LRU

先从简单的东西开始实现。我们的目标是实现一个分布式的、可随节点扩容而扩展的缓存。实际上我们所说的缓存就是一个kv数据库，考虑到缓存不需要严格高可用，而且内存资源宝贵，不可能像mysql主从架构一样奢侈，所以我们只需要想办法把让kv均匀分布在多个节点上即可。很容易想到，我们可以实现一个分布式的hashMap，想办法让分布均匀就好。

作为缓存，我们首先实现一版单机的LRU，使得内存使用超限时能自动淘汰部分数据。此时先不考虑并发问题。

LRU在go里的一个简单的实现方式是一个map+一个linkedlist（container库里的list），list储存具体的kv对，map映射k和对应list_node。每次put或者get一个kv就把那个节点往队尾放，淘汰的时候取队首删除即可。

结构定义：

```go
import "container/list"

type Cache struct {
	maxBytes  int64
	nowBytes  int64
	cache     map[string]*list.Element
	innerList *list.List
	OnEvicted func(key string, value CacheValue)
}

type CacheEntry struct {
	key   string
	value CacheValue
}

type CacheValue interface {
	Len() int
}
```

其中Cache补充了最大字节限制、当前字节占用（通过k、v的len计算）、回调函数（删除一个kv时调用）。这里直接使用container标准库里的链表来替代手动实现，其中list.Element是链表的节点定义，后面使用cache内部字段的时候要记得用反射转成CacheEntry。

之后就是完成对应的功能方法：

```go
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
		return v.Value.(*CacheEntry).value, true
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
	delete(c.cache, oldEntry.key)
	c.nowBytes -= int64(len(oldEntry.key)) + int64(oldEntry.value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(oldEntry.key, oldEntry.value)
	}
}

func (c *Cache) Set(key string, value CacheValue) {
	element, ok := c.cache[key]
	if ok {
		c.innerList.MoveToBack(element)
		entry := element.Value.(*CacheEntry)
		c.nowBytes = c.nowBytes - int64(entry.value.Len()) + int64(value.Len())
		entry.value = value
	} else {
		element = c.innerList.PushBack(&CacheEntry{
			key:   key,
			value: value,
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
```

实现的时候注意下细节即可。最好同步补一些UT确保无bug。

## #0x03 实现单机并发缓存

缓存一定是要支持并发的，这样才能同时照顾多个线程的存取需要，同时保持较高的性能。7daysgolang的实现方法非常粗暴，加一个mutex就结束了，但实际上这样会导致并发的优势完全无法体现。所以这里参考godis，实现了一个分段锁的机制，定义新的数据结构`Cube`作为具名的缓存单机实例，包含若干个shard（也就是第一部分实现的lru缓存）以及每个shard对应的读写锁。此外Cube会存储一个getterFunc，由用户提供，为缓存未命中时用户指定的下层数据获取方法。

```go
// Cube is a concurrency-safe cache instance with a name
type Cube struct {
	shards []*lru.Cache
	mu     []sync.RWMutex
	name   string
	// getterFunc is the custom func to call to get the target value
	getterFunc func(key string) (value lru.CacheValue, err error)
}

func New(name string, getterFunc func(key string) (value lru.CacheValue, err error),
	OnEvicted func(key string, value lru.CacheValue), maxBytes int64) *Cube {
	cube := &Cube{
		shards:     make([]*lru.Cache, 32),
		mu:         make([]sync.RWMutex, 32),
		name:       name,
		getterFunc: getterFunc,
	}
	for i := range cube.shards {
		cube.shards[i] = lru.New(maxBytes/32, OnEvicted)
	}
	return cube
}
```

分段锁的机制实际上就是通过fnv之类的哈希算法，将key归类到不同的shard中，由于每个shard对应一个锁，读写一个shard就不影响其他shard，从而可以提高性能。

```go
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
	if !ok && c.getterFunc != nil {
		valueOutsideCache, err := c.getterFunc(key)
		if err == nil {
			value = valueOutsideCache
			ok = true
			c.shards[shard].Set(key, value)
		}
	}
	return value, ok
}
```

由于我们实现的是分布式缓存而不是数据库，所以不需要支持遍历和持久化等，设计起来比较简单。

## #0x04 实现接入http

系统在设计上除了作为一个单机缓存包被代码引入使用，还应当作为一个单独的分布式中间件被使用，并作彼此之间的沟通协调。这里使用gin框架来完成一个简单的http接入。

```go
func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/cube/:subPath", handleGet)
	r.POST("/cube/:subPath", handlePost)
	return r
}
```

处理get和post请求：

```go
func handleGet(ctx *gin.Context) {
	cubeName := ctx.Param("subPath")
	cube, ok := cubeCache.Cubes[cubeName]
	if !ok {
		ctx.JSON(400, gin.H{"msg": "cube " + cubeName + " not found"})
		return
	}
	key := ctx.Param("key")
	byteValue, ok := cube.Get(key)
	if !ok {
		ctx.JSON(401, gin.H{"msg": "user getter func for " + cubeName + "/" + key + " error"})
		return
	}
	bytes := byteValue.(*lru.Bytes)
	ctx.Data(200, "application/octet-stream", bytes.B)
}

func handlePost(ctx *gin.Context) {
	cubeName := ctx.Param("subPath")
	cube, ok := cubeCache.Cubes[cubeName]
	if !ok {
		ctx.JSON(400, gin.H{"msg": "cube " + cubeName + " not found"})
		return
	}
	key := ctx.Param("key")
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "read request body fail"})
		return
	}
	byteValue := &lru.Bytes{B: body}
	cube.Set(key, byteValue)
	ctx.JSON(200, gin.H{"msg": "success"})
}
```

这里将subPath作为一个参数传递给处理函数，实际上对应的是请求url的最后一个子目录，也就是我们的cube名称。在Post和Get的时候缓存value都以二进制（byte数组）形式存放在body中，参数通过param传递。

## #0x05 实现一致性哈希

下面就实现面向分布式的第一步：一致性哈希。

为什么需要一致性哈希？我们在设计分布式缓存的时候，就需要尽可能保证对于每个key，我们都只让其对应唯一一个节点，避免反复在不同节点拉取数据占用多倍的时间和空间。

一个最理想情况下的方法就是给节点编个号，使用哈希算法把key散布到不同节点。然而对于一个集群，往往涉及节点的动态变化，直接使用哈希来映射key和节点编号，可能导致一旦集群出现细微增删，所有映射就全部失效，导致目前已在本地的所有数据全部无效，致使瞬时DB请求大导致雪崩，这对于缓存来说是不可接受的。

一致性哈希算法解决的问题就是对key在动态节点变化下的散布问题。它的解决方法有点类似时间轮算法，引入一个环，将节点散布在环上，再将key散布到环上。每个key对应的节点就是沿环顺时针找到的第一个节点。这样，在产生节点变动时，影响到的key就只有该节点到前一个节点之间的部分，而非全部。

![](http://cdn.mcyou.cc/20231104170827.png)

当然，考虑到节点相对于环的分布是随机散列的，可能会产生数据不平衡的问题，这里可以通过引入虚拟节点的方式来解决。我们为每个物理节点都设置若干个虚拟节点，再将这些虚拟节点散列在环上，同时维护一个虚拟节点与物理节点之间的映射。这样，通过增加了节点的数量，我们就以很小的代价解决了数据倾斜的问题。

具体代码实现上，我们引入一个mapper作为维护哈希环的数据结构，提供AddNode、RemoveNode和Get方法来为CubeCache提供key-node的映射服务：

```go
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
```

我们使用数组nodes来表示哈希环，使用sort库排序保持其始终递增。在Get时使用sort.SearchInts，这是自带的二分查找方法，来取得key对应的node。

