package prime

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type LruCache struct {
	mu        *sync.Mutex
	maxSize   int
	items     map[interface{}]*list.Element
	evictList *list.List
}

type LruItem struct {
	key   interface{}
	value interface{}
}

func NewLRUCache(maxSize int) *LruCache {
	return &LruCache{
		maxSize:   maxSize,
		mu:        new(sync.Mutex),
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element, maxSize),
	}
}

func (l *LruCache) Get(key interface{}) (value *list.Element, ok bool) {
	defer l.mu.Unlock()
	l.mu.Lock()
	// 最快判断map中是否存在
	if value, ok = l.items[key]; ok {
		// 记录一次使用记录
		l.evictList.MoveToFront(value)
	}
	return
}

func (l *LruCache) Set(key, value interface{}) {
	defer l.mu.Unlock()
	l.mu.Lock()
	// 最快判断map中是否存在
	if val, ok := l.items[key]; ok {
		l.evictList.MoveToFront(val)
		l.items[key] = val
	} else {
		// 判断当前的lruCache是否还存在剩余存储空间
		if l.evictList.Len() < l.maxSize {
			//l.
		} else {
			lastItem := l.evictList.Back()
			fmt.Println("lastItem", lastItem)
			fmt.Println("lastItem.Value", lastItem.Value)
			fmt.Println("lastItem.Value.type", reflect.TypeOf(lastItem.Value))
			lruItem := lastItem.Value.(*LruItem)
			delete(l.items, lruItem.key)
			l.evictList.Remove(lastItem)
		}
		lruItem := &LruItem{
			key:   key,
			value: value,
		}
		l.items[key] = l.evictList.PushFront(lruItem)
	}
}

func TestLru(t *testing.T) {
	lruCache := NewLRUCache(5)
	lruCache.Set(1, 'a')
	lruCache.Set(2, 'b')
	lruCache.Set(3, 'c')
	lruCache.Set(4, 'd')
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
	fmt.Println(lruCache.Get(5))
	lruCache.Set(5, 'd')
	fmt.Println(lruCache.Get(5))
	lruCache.Set(6, 'e')
	val1, ok1 := lruCache.Get(1)
	fmt.Println(val1, ok1)
}
