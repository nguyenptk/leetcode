// https://leetcode.com/problems/lru-cache
package medium

import "container/list"

/*
There are two types of implementation:
	- Build a DoublyListNode
	- Use container/list
*/

// 1. Build a DoublyListNode
//
// type DoublyListNode struct {
// 	key   int
// 	value int
// 	prev  *DoublyListNode
// 	next  *DoublyListNode
// }

// type LRUCache struct {
// 	capacity int
// 	left     *DoublyListNode // LRU
// 	right    *DoublyListNode // most recent
// 	cache    map[int]*DoublyListNode
// }

// func ConstructorLRUCache(capacity int) LRUCache {
// 	left := &DoublyListNode{}
// 	right := &DoublyListNode{}
// 	node := LRUCache{
// 		capacity: capacity,
// 		left:     left,
// 		right:    right,
// 		cache:    make(map[int]*DoublyListNode),
// 	}
// 	node.left.next = right
// 	node.right.prev = left
// 	return node
// }

// func (f *LRUCache) Remove(node *DoublyListNode) {
// 	prev, next := node.prev, node.next
// 	prev.next, next.prev = next, prev
// }

// func (f *LRUCache) Insert(node *DoublyListNode) {
// 	prev, next := f.right.prev, f.right
// 	next.prev = node
// 	prev.next = next.prev
// 	node.prev, node.next = prev, next
// }

// func (f *LRUCache) Get(key int) int {
// 	if _, ok := f.cache[key]; ok {
// 		f.Remove(f.cache[key])
// 		f.Insert(f.cache[key])
// 		return f.cache[key].value
// 	}
// 	return -1
// }

// func (f *LRUCache) Put(key int, value int) {
// 	if _, ok := f.cache[key]; ok {
// 		f.Remove(f.cache[key])
// 	}
// 	f.cache[key] = &DoublyListNode{key: key, value: value}
// 	f.Insert(f.cache[key])
// 	if len(f.cache) > f.capacity {
// 		lru := f.left.next
// 		f.Remove(lru)
// 		delete(f.cache, lru.key)
// 	}
// }

// =========================

// 2. Use container/list

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

type NodeCache struct {
	key   int
	value int
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (f *LRUCache) Get(key int) int {
	if node, ok := f.cache[key]; ok {
		val := node.Value.(*NodeCache).value
		f.list.MoveToBack(node)
		return val
	}
	return -1
}

func (f *LRUCache) Put(key, value int) {
	if node, ok := f.cache[key]; ok {
		node.Value.(*NodeCache).value = value
		f.list.MoveToBack(node)
		return
	}
	node := &NodeCache{key: key, value: value}
	data := f.list.PushBack(node)
	f.cache[key] = data
	if f.list.Len() > f.capacity {
		idx := f.list.Front().Value.(*NodeCache).key
		delete(f.cache, idx)
		f.list.Remove(f.list.Front())
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
