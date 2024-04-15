// https://leetcode.com/problems/clone-graph
package medium

type NodeGraph struct {
	Val       int
	Neighbors []*NodeGraph
}

func CloneGraph(nodeGraph *NodeGraph) *NodeGraph {
	return dfsCloneGraph(nodeGraph, map[*NodeGraph]*NodeGraph{})
}

func dfsCloneGraph(nodeGraph *NodeGraph, visited map[*NodeGraph]*NodeGraph) *NodeGraph {
	if nodeGraph == nil {
		return nil
	}
	if _, ok := visited[nodeGraph]; ok {
		return visited[nodeGraph]
	}
	clone := &NodeGraph{
		Val: nodeGraph.Val,
	}
	visited[nodeGraph] = clone
	for _, n := range nodeGraph.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfsCloneGraph(n, visited))
	}
	return clone
}
