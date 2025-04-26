// https://leetcode.com/problems/design-hashmap/
package easy

type Pair struct {
	key   int
	value int
	next  *Pair
}

// MyHashMap represents the hashmap using separate chaining with linked lists
type MyHashMap struct {
	buckets []*Pair
	size    int
}

// Constructor initializes the hashmap
func ConstructorMyHashMap() MyHashMap {
	size := 1000
	buckets := make([]*Pair, size)
	return MyHashMap{buckets: buckets, size: size}
}

// Hash function to determine bucket index
func (f *MyHashMap) hash(key int) int {
	return key % f.size
}

// Put inserts or updates a key-value pair
func (f *MyHashMap) Put(key int, value int) {
	idx := f.hash(key)
	head := f.buckets[idx]

	// Traverse the linked list to update existing key
	curr := head
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}

	// Insert new key-value pair at the beginning
	newPair := &Pair{key: key, value: value, next: head}
	f.buckets[idx] = newPair
}

// Get retrieves the value for a key
func (f *MyHashMap) Get(key int) int {
	idx := f.hash(key)
	curr := f.buckets[idx]

	// Traverse the linked list to find the key
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}

	return -1 // Key not found
}

// Remove deletes a key-value pair from the hashmap
func (f *MyHashMap) Remove(key int) {
	idx := f.hash(key)
	curr := f.buckets[idx]
	var prev *Pair

	// Traverse to find the key
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				// Removing head node
				f.buckets[idx] = curr.next
			} else {
				// Bypass the current node
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
