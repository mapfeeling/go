package classics

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

type LRUCache struct {
	mu        sync.Mutex
	size      int
	items     map[interface{}]*list.Element
	evictList *list.List
}

type LRUItem struct {
	key   interface{}
	value interface{}
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		size:      maxSize,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element, maxSize),
	}
}

func (lru *LRUCache) Get(key interface{}) (interface{}, bool) {
	var value interface{}
	if element, ok := lru.items[key]; ok {
		lru.evictList.MoveToFront(element)
		value = element.Value.(*LRUItem).value
		return value, true
	}
	return value, false
}

func (lru *LRUCache) RemoveElement() {
	element := lru.evictList.Back()
	lru.evictList.Remove(element)
	entry := element.Value.(*LRUItem)
	delete(lru.items, entry.key)
}

func (lru *LRUCache) Put(key interface{}, value interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	var item *LRUItem
	if element, ok := lru.items[key]; ok {
		lru.evictList.MoveToFront(element)
		item = element.Value.(*LRUItem)
		item.value = value
	} else {
		if lru.evictList.Len() >= lru.size {
			lru.RemoveElement()
		}
		item = &LRUItem{
			key:   key,
			value: value,
		}
		lru.items[key] = lru.evictList.PushFront(item)
	}
}

func TestMyLRUCache(t *testing.T) {
	lruCache := NewLRUCache(3)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
	fmt.Println(lruCache.Get(3))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(a())
}

func a() (value interface{}, okk bool) {
	var m = make(map[interface{}]interface{}, 1)
	m[1] = 1
	if value, okk = m[1]; okk {
		fmt.Println(value)
	}
	return
}
