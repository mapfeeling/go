package minor

import (
	"fmt"
	"testing"
)

type LinkNode struct {
	key, value interface{}
	prev, next *LinkNode
}

type LRUCache struct {
	hashMap    map[interface{}]*LinkNode
	capacity   int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{struct{}{}, struct{}{}, nil, nil}
	tail := &LinkNode{struct{}{}, struct{}{}, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[interface{}]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.prev = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Get(key interface{}) interface{} {
	hashMap := this.hashMap
	if node, ok := hashMap[key]; ok {
		this.MoveToHead(node)
		return node
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key interface{}, value interface{}) {
	hashMap := this.hashMap
	if node, ok := hashMap[key]; ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		n := &LinkNode{key, value, nil, nil}
		if len(hashMap) >= this.capacity {
			delete(hashMap, this.tail.prev.key)
			this.RemoveNode(this.tail.prev)
		}
		hashMap[key] = n
		this.AddNode(n)
	}
}

func TestLRU(t *testing.T) {
	lruCache := Constructor(10)
	fmt.Println(lruCache.Get("1"))
	lruCache.Put("1", 1)
	println(lruCache.Get("1"), 12)
}
