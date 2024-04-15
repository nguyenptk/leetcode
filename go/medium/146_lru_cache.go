// https://leetcode.com/problems/lru-cache
package medium

type LRUCache struct {
	left     *DoublyListNode
	right    *DoublyListNode
	cache    map[int]*DoublyListNode
	capacity int
}

type DoublyListNode struct {
	key   int
	value int
	next  *DoublyListNode
	prev  *DoublyListNode
}

func ConstructorLRUCache(capacity int) LRUCache {
	ret := LRUCache{
		left:     &DoublyListNode{key: 0, value: 0},
		right:    &DoublyListNode{key: 0, value: 0},
		cache:    make(map[int]*DoublyListNode),
		capacity: capacity,
	}
	ret.left.next = ret.right
	ret.right.prev = ret.left
	return ret
}

func (f *LRUCache) Remove(node *DoublyListNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (f *LRUCache) Insert(node *DoublyListNode) {
	prev := f.right.prev
	next := f.right
	next.prev = node
	prev.next = next.prev
	node.next = next
	node.prev = prev
}

func (f *LRUCache) Get(key int) int {
	if _, ok := f.cache[key]; ok {
		f.Remove(f.cache[key])
		f.Insert(f.cache[key])
		return f.cache[key].value
	}
	return -1
}

func (f *LRUCache) Put(key int, value int) {
	if _, ok := f.cache[key]; ok {
		f.Remove(f.cache[key])
	}
	f.cache[key] = &DoublyListNode{key: key, value: value}
	f.Insert(f.cache[key])

	if len(f.cache) > f.capacity {
		// remove from the list and delete the LRU from hashmap
		lru := f.left.next
		f.Remove(lru)
		delete(f.cache, lru.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
