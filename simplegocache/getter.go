package simplegocache

import (
	"sync"
)

// Getter 为键加载数据
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc 实现Getter的Get方法
type GetterFunc func(key string) ([]byte, error)

// Get 实现Getter
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)
