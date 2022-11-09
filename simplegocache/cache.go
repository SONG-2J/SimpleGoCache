package simplegocache

import (
	"SimpleGoCache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex // 互斥锁
	lru        *lru.Cache // lru缓存
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil { // 延迟初始化，判断了c.lru是否为nil，如果等于nil再创建实例
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}
