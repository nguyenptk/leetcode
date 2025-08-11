// https://leetcode.com/problems/course-schedule
package medium

// Topology sort + bfs
func CanFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		from := p[0]
		to := p[1]
		adj[to] = append(adj[to], from)
		inDegrees[from]++
	}

	queue := make([]int, 0)
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited++

		for _, nei := range adj[curr] {
			inDegrees[nei]--
			if inDegrees[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	return numCourses == visited
}

// Time complexity: O(m+n)
// Space complexity: O(m+n)

// func canFinish(numCourses int, prerequisites [][]int) bool {
//     adj := make(map[int][]int)
//     for _, pre := range prerequisites {
//         from := pre[0]
//         to := pre[1]
//         adj[to] = append(adj[to], from)
//     }

//     visit := make([]bool, numCourses)
//     inStack := make([]bool, numCourses)

//     var dfs func(node int) bool
//     dfs = func(node int) bool {
//         if inStack[node] {
//             return true
//         }
//         if visit[node] {
//             return false
//         }

//         inStack[node] = true
//         visit[node] = true
//         for _, nei := range adj[node] {
//             if dfs(nei) {
//                 return true
//             }
//         }

//         inStack[node] = false
//         return false
//     }

//     for i := range numCourses {
//         if dfs(i) {
//             return false
//         }
//     }

//     return true
// }
