// https://leetcode.com/problems/lfu-cache/
package hard

type DoublyListNode struct {
	key   int
	value int
	freq  int
	prev  *DoublyListNode
	next  *DoublyListNode
}

type DoublyLinkedList struct {
	head *DoublyListNode
	tail *DoublyListNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &DoublyListNode{}
	tail := &DoublyListNode{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (dll *DoublyLinkedList) AddNode(node *DoublyListNode) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
	dll.size++
}

func (dll *DoublyLinkedList) RemoveNode(node *DoublyListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	dll.size--
}

func (dll *DoublyLinkedList) PopTail() *DoublyListNode {
	if dll.size == 0 {
		return nil
	}
	node := dll.tail.prev
	dll.RemoveNode(node)
	return node
}

type LFUCache struct {
	capacity int
	size     int
	minFreq  int
	cache    map[int]*DoublyListNode
	freqMap  map[int]*DoublyLinkedList
}

func ConstructorLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*DoublyListNode),
		freqMap:  make(map[int]*DoublyLinkedList),
	}
}

func (cache *LFUCache) Get(key int) int {
	if node, found := cache.cache[key]; found {
		cache.incrementFrequency(node)
		return node.value
	}
	return -1
}

func (cache *LFUCache) Put(key, value int) {
	if cache.capacity == 0 {
		return
	}

	if node, found := cache.cache[key]; found {
		node.value = value
		cache.incrementFrequency(node)
		return
	}

	if cache.size == cache.capacity {
		cache.evict()
	}

	newNode := &DoublyListNode{
		key:   key,
		value: value,
		freq:  1,
	}
	cache.cache[key] = newNode
	if cache.freqMap[1] == nil {
		cache.freqMap[1] = NewDoublyLinkedList()
	}
	cache.freqMap[1].AddNode(newNode)
	cache.minFreq = 1
	cache.size++
}

func (cache *LFUCache) evict() {
	list := cache.freqMap[cache.minFreq]
	node := list.PopTail()

	delete(cache.cache, node.key)
	cache.size--

	if list.size == 0 {
		delete(cache.freqMap, cache.minFreq)
	}
}

func (cache *LFUCache) incrementFrequency(node *DoublyListNode) {
	curFreq := node.freq

	cache.freqMap[curFreq].RemoveNode(node)

	if cache.freqMap[curFreq].size == 0 {
		delete(cache.freqMap, curFreq)
		if cache.minFreq == curFreq {
			cache.minFreq++
		}
	}

	node.freq++
	if cache.freqMap[node.freq] == nil {
		cache.freqMap[node.freq] = NewDoublyLinkedList()
	}
	cache.freqMap[node.freq].AddNode(node)
}
