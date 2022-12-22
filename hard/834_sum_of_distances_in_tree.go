// https://leetcode.com/problems/sum-of-distances-in-tree
package hard

func SumOfDistancesInTree(n int, edges [][]int) []int {
	result := make([]int, n)
	count := make([]int, n)
	for i := 0; i < len(count); i++ {
		count[i] = 1
	}
	tree := make(map[int][]int, n)

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	postorder(tree, 0, -1, count, result)
	preorder(tree, 0, -1, count, result)
	return result
}

func postorder(tree map[int][]int, node, parent int, count []int, result []int) {
	for _, child := range tree[node] {
		if child == parent {
			continue
		}
		postorder(tree, child, node, count, result)
		count[node] += count[child]
		result[node] += result[child] + count[child]
	}
}

func preorder(tree map[int][]int, node, parent int, count []int, result []int) {
	for _, child := range tree[node] {
		if child == parent {
			continue
		}
		result[child] = result[node] - count[child] + (len(tree) - count[child])
		preorder(tree, child, node, count, result)
	}
}
