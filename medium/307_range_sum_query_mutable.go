// https://leetcode.com/problems/range-sum-query-mutable/
package medium

type NumArray struct {
	Nums []int
}

func ConstructorNumArray(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.Nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.Nums[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
