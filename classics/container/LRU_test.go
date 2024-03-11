package container

import (
	"container/list"
	"fmt"
	"sync"
	"testing"

	"github.com/bluele/gcache"
)

type LRUCache struct {
	mu          sync.Mutex
	size        int // 允许使用的最大内存
	evictList   *list.List
	items       map[interface{}]*list.Element
	evictedFunc EvictedFunc // 某条记录被移除时的回调函数
}

type (
	EvictedFunc func(interface{}, interface{})
)

type lruItem struct {
	key   interface{}
	value interface{}
}

func NewLRUCache(maxBytes int, OnEvicted func(interface{}, interface{})) *LRUCache {
	return &LRUCache{
		size:        maxBytes,
		evictList:   list.New(),
		items:       make(map[interface{}]*list.Element),
		evictedFunc: OnEvicted,
	}
}

// Get LRU的查找,第一步从字典中找到对应的双向链表的节点,第二步,将该节点移动到链表尾
func (lru *LRUCache) Get(key string) (kv interface{}, ok bool) {
	// 如果找到节点,将缓存移动到队尾
	if element, ok := lru.items[key]; ok {
		lru.evictList.MoveToBack(element)
		kv = element.Value.(*lruItem)
		return kv, true
	}
	return
}

// RemoveOldest LRU的删除,实际上是缓存淘汰,即移除最近访问的节点,队首
func (lru *LRUCache) removeElement(e *list.Element) {
	lru.evictList.Remove(e)
	entry := e.Value.(*lruItem)
	delete(lru.items, entry.key)
	if lru.evictedFunc != nil {
		entry := e.Value.(*lruItem)
		lru.evictedFunc(entry.key, entry.value)
	}
}

// Put ：LRU的新增or修改
func (lru *LRUCache) Put(key, value interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	var item *lruItem
	// map的key存在,直接修改map的key和value
	if element, ok := lru.items[key]; ok {
		lru.evictList.MoveToBack(element)
		item = element.Value.(*lruItem)
		item.value = value
	} else {
		// map的key不存在,直接修改map的key和value
		// 如果当前已经使用的双链表长度>=lru的最大长度使用淘汰策略
		fmt.Println(value, lru.evictList.Len())
		if lru.evictList.Len() >= lru.size {
			lru.evict(1)
		}
		item = &lruItem{
			key:   key,
			value: value,
		}
		lru.items[key] = lru.evictList.PushFront(item)
		fmt.Println("......", lru.items[key])

	}
}

func (lru *LRUCache) evict(count int) {
	for i := 0; i < count; i++ {
		ent := lru.evictList.Back()
		if ent == nil {
			return
		} else {
			fmt.Println("执行了evict")
			lru.removeElement(ent)
		}
	}
}

func TestLRU(t *testing.T) {
	gc := gcache.New(20).LRU().Build()
	_ = gc.Set("a", "ok")

	value, err := gc.Get("a")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	lru := NewLRUCache(3, nil)
	lru.Put("a", 1)
	lru.Put("b", 2)
	lru.Put("c", 3)
	lru.Put("d", 4)
	lru.Put("e", 5)
	kv, ok := lru.Get("a")
	fmt.Println(kv, ok)
	fmt.Println("a")
	fmt.Println(lru.Get("a"))
	fmt.Println("b")
	fmt.Println(lru.Get("b"))
	fmt.Println("c")
	fmt.Println(lru.Get("c"))
	fmt.Println("d")
	fmt.Println(lru.Get("d"))
	fmt.Println("e")
	fmt.Println(lru.Get("e"))
}
