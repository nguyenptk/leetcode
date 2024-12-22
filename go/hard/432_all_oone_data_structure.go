// https://leetcode.com/problems/all-oone-data-structure
package hard

import "math"

type Node struct {
	key   string
	count int
	next  *Node
	prev  *Node
}

type AllOne struct {
	head    *Node
	tail    *Node
	entries map[string]*Node
}

func ConstructorAllOne() AllOne {
	head := &Node{count: math.MaxInt}
	tail := &Node{count: 0}

	head.prev = tail
	tail.next = head

	return AllOne{
		head:    head,
		tail:    tail,
		entries: make(map[string]*Node),
	}
}

func (f *AllOne) Inc(key string) {
	if f.entries[key] == nil {
		value := &Node{
			key:   key,
			count: 1,
			prev:  f.tail,
			next:  f.tail.next,
		}

		f.tail.next.prev = value
		f.tail.next = value

		f.entries[key] = value
		return
	}

	f.entries[key].count++
	ptr := f.entries[key]
	for ptr.next != nil && ptr.count > ptr.next.count {
		f.swap(ptr, ptr.next)
	}
}

func (f *AllOne) Dec(key string) {
	f.entries[key].count--

	if f.entries[key].count == 0 {
		f.remove(key)
		return
	}

	ptr := f.entries[key]
	for ptr.prev != nil && ptr.count < ptr.prev.count {
		f.swap(ptr.prev, ptr)
	}
}

func (f *AllOne) GetMaxKey() string {
	return f.head.prev.key
}

func (f *AllOne) GetMinKey() string {
	return f.tail.next.key
}

func (f *AllOne) swap(x, y *Node) {
	x.next = y.next
	y.prev = x.prev
	x.prev.next = y
	y.next.prev = x
	x.prev = y
	y.next = x
}

func (f *AllOne) remove(key string) {
	node := f.entries[key]

	node.prev.next = node.next
	node.next.prev = node.prev

	delete(f.entries, key)
}
