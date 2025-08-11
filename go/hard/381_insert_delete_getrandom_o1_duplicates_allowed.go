// https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/
package hard

import "math/rand"

type RandomizedCollection struct {
	nums []int
	idx  map[int][]int // val -> list of positions in nums
	pos  map[int]int   // index in nums -> index in idx[val]
}

func ConstructorRandomizedCollection() RandomizedCollection {
	return RandomizedCollection{
		nums: []int{},
		idx:  make(map[int][]int),
		pos:  make(map[int]int),
	}
}

func (f *RandomizedCollection) Insert(val int) bool {
	f.nums = append(f.nums, val)
	f.idx[val] = append(f.idx[val], len(f.nums)-1)
	f.pos[len(f.nums)-1] = len(f.idx[val]) - 1
	return len(f.idx[val]) == 1
}

func (f *RandomizedCollection) Remove(val int) bool {
	positions, ok := f.idx[val]
	if !ok || len(positions) == 0 {
		return false
	}

	// Get the index in nums to remove
	removeIdx := f.idx[val][len(f.idx[val])-1]
	f.idx[val] = f.idx[val][:len(f.idx[val])-1]

	// Swap with the last element if not the last one
	lastIdx := len(f.nums) - 1
	lastVal := f.nums[lastIdx]

	f.nums[removeIdx] = lastVal
	f.nums = f.nums[:lastIdx]

	if removeIdx != lastIdx {
		// Update idx and pos for lastVal
		posIdx := f.pos[lastIdx]
		f.idx[lastVal][posIdx] = removeIdx
		f.pos[removeIdx] = posIdx
	}

	delete(f.pos, lastIdx)

	if len(f.idx[val]) == 0 {
		delete(f.idx, val)
	}

	return true
}

func (f *RandomizedCollection) GetRandom() int {
	return f.nums[rand.Intn(len(f.nums))]
}

// Time complexity:
// Insert(val int): O(1)
// Remove(val int): O(1)
// GetRandom(): O(1)
// Space complexity:
// nums holds up to n elements → O(n)
// idx[val] can store multiple indices per value → O(n) in worst case
// pos[i] maps indices to slices in idx[val] → O(n)
