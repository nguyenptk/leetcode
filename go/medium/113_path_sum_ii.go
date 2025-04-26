// https://leetcode.com/problems/path-sum-ii/
package medium

func PathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	dfsPathSum(root, targetSum, []int{}, &result)
	return result
}

func dfsPathSum(root *TreeNode, targetSum int, path []int, result *[][]int) {
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		*result = append(*result, append(path, root.Val))
		return
	}
	path = append(append([]int{}, path...), root.Val)
	if root.Left != nil {
		dfsPathSum(root.Left, targetSum-root.Val, path, result)
	}
	if root.Right != nil {
		dfsPathSum(root.Right, targetSum-root.Val, path, result)
	}
}

// func PathSum(root *TreeNode, targetSum int) [][]int {
//     result := make([][]int, 0)

//     var dfs func(root *TreeNode, targetSum int, path []int)
//     dfs = func(root *TreeNode, targetSum int, path []int) {
//         if root == nil {
//             return
//         }

//         subSum := targetSum - root.Val
//         path = append(path, root.Val)
//         if subSum == 0 && root.Left == nil && root.Right == nil {
//             copiedPath := make([]int, len(path))
//             copy(copiedPath, path)
//             result = append(result, copiedPath)
//         }

//         dfs(root.Left, subSum, path)
//         dfs(root.Right, subSum, path)

//         path = path[:len(path)-1]
//     }

//     path := make([]int, 0)
//     dfs(root, targetSum, path)

//     return result
// }
