// https://leetcode.com/problems/clone-graph
package medium

type NodeGraph struct {
	Val       int
	Neighbors []*NodeGraph
}

func CloneGraph(nodeGraph *NodeGraph) *NodeGraph {
	visit := make(map[*NodeGraph]*NodeGraph)

	var dfs func(node *NodeGraph) *NodeGraph
	dfs = func(node *NodeGraph) *NodeGraph {
		if node == nil {
			return nil
		}

		if curr, ok := visit[node]; ok {
			return curr
		}
		clone := &NodeGraph{
			Val: node.Val,
		}
		visit[node] = clone

		for _, nei := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(nei))

		}

		return clone
	}

	return dfs(nodeGraph)

}
