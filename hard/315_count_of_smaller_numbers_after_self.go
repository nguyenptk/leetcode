// https://leetcode.com/problems/counter-of-smaller-numbers-after-self/
package hard

import "sort"

type BinaryIndexedTree struct {
	Values []int
}

func initBinaryIndexedTree(len int) *BinaryIndexedTree {
	return &BinaryIndexedTree{
		Values: make([]int, len),
	}
}

// Update the current key, value from map to BinaryIndexedTree
func (b *BinaryIndexedTree) update(k, v int) {
	for ; k < len(b.Values); k += (k & (-k)) {
		b.Values[k] += v
	}
}

// Sum the values that are les than the current value
func (b *BinaryIndexedTree) sum(r int) int {
	sum := 0
	for ; r > 0; r -= (r & (-r)) {
		sum += b.Values[r]
	}
	return sum
}

func CountSmaller(nums []int) []int {
	n := len(nums)

	// Init BinaryIndexedTree
	bit := initBinaryIndexedTree(n)

	// Create a new temp array to store the ascending
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = nums[i]
	}
	sort.Ints(temp)

	// Create a map to store the unique value
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[temp[i]]; !ok {
			mp[temp[i]] = i + 1
		}
	}

	result := make([]int, n)
	// Iterate the nums array to find the result
	for i := n - 1; i >= 0; i-- {
		result[i] = bit.sum(mp[nums[i]] - 1)
		bit.update(mp[nums[i]], 1)
	}
	return result
}
