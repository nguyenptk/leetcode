// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return recSortedArrayToBST(nums, 0, len(nums)-1)
}

func recSortedArrayToBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = recSortedArrayToBST(nums, start, mid-1)
	root.Right = recSortedArrayToBST(nums, mid+1, end)

	return root
}
