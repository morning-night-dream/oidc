package cache

import (
	"fmt"
	"sync"
)

type Cache[T any] struct {
	lock sync.Mutex
	data map[string]T
}

func (cc *Cache[T]) Get(key string) (T, error) {
	var val T

	cc.lock.Lock()
	defer cc.lock.Unlock()

	if _, ok := cc.data[key]; !ok {
		return val, fmt.Errorf("key %s not found", key)
	}

	return cc.data[key], nil
}

func (cc *Cache[T]) Set(key string, val T) error {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	cc.data[key] = val

	return nil
}
