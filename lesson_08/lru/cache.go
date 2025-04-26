package lru

import (
	"container/list"
)

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type CacheStorage struct {
	items     map[string]*list.Element
	capacity  int
	cacheList *list.List
}

func NewLruCache(capacity int) LruCache {
	cacheStorage := CacheStorage{
		items:     make(map[string]*list.Element, capacity),
		capacity:  capacity,
		cacheList: list.New(),
	}

	return &cacheStorage
}

type cacheItem struct {
	key   string
	value string
}

func (c *CacheStorage) Put(key, value string) {
	if elem, exists := c.items[key]; exists {
		c.cacheList.MoveToFront(elem)
		elem.Value = &cacheItem{
			key:   key,
			value: value,
		}

		return
	}

	if len(c.items) >= c.capacity {
		last := c.cacheList.Back()
		if last != nil {
			item := last.Value.(*cacheItem)
			delete(c.items, item.key)
			c.cacheList.Remove(last)
		}
	}

	item := cacheItem{
		key:   key,
		value: value,
	}

	elem := c.cacheList.PushFront(&item)
	c.items[key] = elem
}

func (c *CacheStorage) Get(key string) (string, bool) {
	elem, exists := c.items[key]
	if !exists {
		return "", false
	}

	c.cacheList.MoveToFront(elem)

	return elem.Value.(*cacheItem).value, true
}
