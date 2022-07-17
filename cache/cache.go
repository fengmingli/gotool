package cache

/**
 * @Author: LFM
 * @Date: 2022/7/17 16:45
 * @Since: 1.0.0
 * @Desc: TODO
 */

import (
	"sync"
)

const (
	shardsCount int = 32
)

type Cache []*cacheShard

func NewCache(maxSize int) Cache {
	if maxSize < shardsCount {
		maxSize = shardsCount
	}
	cache := make(Cache, shardsCount)
	for i := 0; i < shardsCount; i++ {
		cache[i] = &cacheShard{
			items:   make(map[uint64]interface{}),
			maxSize: maxSize / shardsCount,
		}
	}
	return cache
}

func (c Cache) getShard(index uint64) *cacheShard {
	return c[index%uint64(shardsCount)]
}

// Add Returns true if object already existed, false otherwise.
func (c *Cache) Add(index uint64, obj interface{}) bool {
	return c.getShard(index).add(index, obj)
}

func (c *Cache) Get(index uint64) (obj interface{}, found bool) {
	return c.getShard(index).get(index)
}

type cacheShard struct {
	items map[uint64]interface{}
	sync.RWMutex
	maxSize int
}

// Returns true if object already existed, false otherwise.
func (s *cacheShard) add(index uint64, obj interface{}) bool {
	s.Lock()
	defer s.Unlock()
	_, isOverwrite := s.items[index]
	if !isOverwrite && len(s.items) >= s.maxSize {
		var randomKey uint64
		for randomKey = range s.items {
			break
		}
		delete(s.items, randomKey)
	}
	s.items[index] = obj
	return isOverwrite
}

func (s *cacheShard) get(index uint64) (obj interface{}, found bool) {
	s.RLock()
	defer s.RUnlock()
	obj, found = s.items[index]
	return
}
