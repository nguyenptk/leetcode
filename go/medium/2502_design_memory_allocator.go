// https://leetcode.com/problems/design-memory-allocator/
package medium

type Allocator struct {
	memory []int
}

func ConstructorAllocator(n int) Allocator {
	return Allocator{
		memory: make([]int, n),
	}
}

func (f *Allocator) Allocate(size int, mID int) int {
	var count int = 0
	var start int = 0
	for i := 0; i < len(f.memory); i++ {
		if f.memory[i] == 0 {
			count++
			if count == size {
				for j := 0; j < size; j++ {
					f.memory[j+start] = mID
				}
				return start
			}
		} else {
			start = i + 1
			count = 0
		}
	}
	return -1
}

func (f *Allocator) FreeMemory(mID int) int {
	var count int = 0
	for i := 0; i < len(f.memory); i++ {
		if f.memory[i] == mID {
			f.memory[i] = 0
			count++
		}
	}
	return count
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
