package cache

/**
 * @Author: LFM
 * @Date: 2022/12/11 16:31
 * @Since: 1.0.0
 * @Desc: 多层map线程安全操作
 */

import (
	"sync"
)

type MultiMap struct {
	data map[string]map[string]string
	mu   sync.RWMutex
}

func (mm *MultiMap) Put(key1, key2 string, value string) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	if mm.data == nil {
		mm.data = make(map[string]map[string]string)
	}
	if mm.data[key1] == nil {
		mm.data[key1] = make(map[string]string)
	}
	mm.data[key1][key2] = value
}

func (mm *MultiMap) Get(key1, key2 string) (value string, ok bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if mm.data == nil {
		return "", false
	}
	m, ok := mm.data[key1]
	if !ok {
		return "", false
	}
	value, ok = m[key2]
	return value, ok
}

func (mm *MultiMap) Delete(key1, key2 string) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	delete(mm.data[key1], key2)
}
