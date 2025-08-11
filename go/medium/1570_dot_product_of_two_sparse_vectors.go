// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/
package medium

type SparseVector struct {
	mapping map[int]int
}

func ConstructorSparseVector(nums []int) SparseVector {
	mapping := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			mapping[i] = nums[i]
		}
	}
	return SparseVector{
		mapping: mapping,
	}
}

// Return the dotProduct of two sparse vectors
func (f *SparseVector) dotProduct(vec SparseVector) int {
	result := 0
	for k, v := range f.mapping {
		if val, ok := vec.mapping[k]; ok {
			result += v * val
		}
	}

	return result
}

// Time complexity: O(n) for creating the Hash Map; O(L) for calculating the dot product.
// Space complexity: O(L) for creating the Hash Map, as we only store elements that are non-zero. O(1) for calculating the dot product.

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
