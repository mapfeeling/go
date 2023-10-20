package classics

import (
	"container/list"
	"log"
	"testing"
)

type LRUCache struct {
	maxBytes  int64 // 允许使用的最大内存
	nBytes    int64 // 当前已经使用的内存
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func NewLRUCache(maxBytes int64, OnEvicted func(string, Value)) *LRUCache {
	return &LRUCache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

func (lru *LRUCache) Get(key string) (val Value, ok bool) {
	// 如果找到节点,将缓存移动到队尾
	if element, ok := lru.cache[key]; ok {
		lru.ll.MoveToBack(element)
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

func (lru *LRUCache) RemoveOldest() {
	element := lru.ll.Front() // 取到首节点,从链表中删除
	if element != nil {
		lru.ll.Remove(element)
		kv := element.Value.(*entry)
		delete(lru.cache, kv.key)
		lru.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if lru.OnEvicted != nil {
			lru.OnEvicted(kv.key, kv.value)
		}
	}
}

func (lru *LRUCache) Put(key string, value Value) {
	if int64(value.Len()) > lru.maxBytes {
		log.Printf("超过最大容量%d,插入缓存容量为%d", lru.maxBytes, int64(value.Len()))
		return
	}
	if element, ok := lru.cache[key]; ok {
		lru.ll.MoveToBack(element)
		kv := element.Value.(*entry)
		lru.nBytes += int64(value.Len()) - int64(kv.value.Len())
	} else {

	}
}

func TestLRU(t *testing.T) {

}
