// https://leetcode.com/problems/boundary-of-binary-tree/
package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BoundaryOfBinaryTree(root *TreeNode) []int {
	leftBoundary := []int{}
	rightBoundary := []int{}
	leaves := []int{}

	preorder(root, &leftBoundary, &rightBoundary, &leaves, 0)

	result := append(leftBoundary, leaves...)
	result = append(result, rightBoundary...)

	return result
}

func preorder(
	node *TreeNode,
	leftBoundary *[]int,
	rightBoundary *[]int,
	leaves *[]int,
	flag int,
) {
	if node == nil {
		return
	}
	if isRightBoundary(flag) {
		*rightBoundary = append([]int{node.Val}, *rightBoundary...)
	} else if isLeftBoundary(flag) || isRoot(flag) {
		*leftBoundary = append(*leftBoundary, node.Val)
	} else if isLeaf(node) {
		*leaves = append(*leaves, node.Val)
	}
	preorder(node.Left, leftBoundary, rightBoundary, leaves, leftChildFlag(node, flag))
	preorder(node.Right, leftBoundary, rightBoundary, leaves, rightChildFlag(node, flag))
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func isLeftBoundary(flag int) bool {
	return flag == 1
}

func isRightBoundary(flag int) bool {
	return flag == 2
}

func isRoot(flag int) bool {
	return flag == 0
}

func leftChildFlag(node *TreeNode, flag int) int {
	if isLeftBoundary(flag) || isRoot(flag) {
		return 1
	} else if isRightBoundary(flag) && node.Right == nil {
		return 2
	} else {
		return 3
	}
}

func rightChildFlag(node *TreeNode, flag int) int {
	if isRightBoundary(flag) || isRoot(flag) {
		return 2
	} else if isLeftBoundary(flag) && node.Left == nil {
		return 1
	} else {
		return 3
	}
}

// Time complexity: O(n) One complete traversal of the tree is done.
// Space complexity: O(n) The recursive stack can grow upto a depth of n. Further, left_boundary, right_boundary and leaves combined together can be of size n.
