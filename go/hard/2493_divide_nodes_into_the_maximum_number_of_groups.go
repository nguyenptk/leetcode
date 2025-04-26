// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/
package hard

import "container/list"

// Find the maximum number of groups in the graph
func MagnificentSets(n int, edges [][]int) int {
	adjList := make([][]int, n)
	parent := make([]int, n)
	depth := make([]int, n)

	// Initialize parent array for Union-Find
	for i := range parent {
		parent[i] = -1
	}

	// Build adjacency list and apply Union-Find
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
		unionRec(u, v, parent, depth)
	}

	numOfGroupsForComponent := make(map[int]int)

	// Calculate max number of groups for each connected component
	for node := 0; node < n; node++ {
		numberOfGroups := getNumberOfGroups(adjList, node, n)
		if numberOfGroups == -1 {
			return -1 // Invalid partition
		}
		rootNode := findRec(node, parent)
		numOfGroupsForComponent[rootNode] = max(numOfGroupsForComponent[rootNode], numberOfGroups)
	}

	// Sum all max groups for each component
	totalGroups := 0
	for _, val := range numOfGroupsForComponent {
		totalGroups += val
	}

	return totalGroups
}

// Perform BFS to calculate the number of groups (layers) in a component
func getNumberOfGroups(adjList [][]int, srcNode, n int) int {
	queue := list.New()
	layerSeen := make([]int, n)
	for i := range layerSeen {
		layerSeen[i] = -1
	}

	queue.PushBack(srcNode)
	layerSeen[srcNode] = 0
	deepestLayer := 0

	for queue.Len() > 0 {
		numOfNodesInLayer := queue.Len()
		for i := 0; i < numOfNodesInLayer; i++ {
			currentNode := queue.Remove(queue.Front()).(int)
			for _, neighbor := range adjList[currentNode] {
				if layerSeen[neighbor] == -1 {
					layerSeen[neighbor] = deepestLayer + 1
					queue.PushBack(neighbor)
				} else if layerSeen[neighbor] == deepestLayer {
					return -1 // Invalid partition (cycle detection)
				}
			}
		}
		deepestLayer++
	}

	return deepestLayer
}

// Find the root of a node using Union-Find with path compression
func findRec(node int, parent []int) int {
	if parent[node] == -1 {
		return node
	}
	parent[node] = findRec(parent[node], parent)
	return parent[node]
}

// Union two sets by rank (depth)
func unionRec(node1, node2 int, parent, depth []int) {
	root1 := findRec(node1, parent)
	root2 := findRec(node2, parent)

	if root1 == root2 {
		return
	}

	if depth[root1] < depth[root2] {
		root1, root2 = root2, root1
	}

	parent[root2] = root1
	if depth[root1] == depth[root2] {
		depth[root1]++
	}
}
