// https://leetcode.com/problems/serialize-and-deserialize-binary-tree

package hard

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

// func Constructor() Codec {
// 	return Codec{}
// }

// Serializes a tree to a single string.
func (f *Codec) Serialize(root *TreeNode) string {
	result := make([]string, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = append(result, "N")
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (f *Codec) Deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if vals[i] == "N" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(vals[i])
		node := &TreeNode{Val: val}
		i++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
