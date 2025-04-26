// https://leetcode.com/problems/course-schedule
package medium

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, 0)
	for i := 0; i < numCourses; i++ {
		graph = append(graph, make([]int, 0))
	}
	for _, p := range prerequisites {
		from := p[1]
		to := p[0]
		graph[to] = append(graph[to], from)
	}

	result := make([]int, 0)
	visit := make([]bool, numCourses)
	cycle := make([]bool, numCourses)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if cycle[course] {
			return false
		} else if visit[course] {
			return true
		}
		cycle[course] = true
		for _, p := range graph[course] {
			if !dfs(p) {
				return false
			}
		}
		cycle[course] = false
		visit[course] = true
		result = append(result, course)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return result
}

// Topology Sort approach
// func FindOrder(numCourses int, prerequisites [][]int) []int {
//     result := make([]int, 0)

//     adj := make(map[int][]int)
//     inDegrees := make([]int, numCourses)
//     for _, pre := range prerequisites {
//         from := pre[0]
//         to := pre[1]
//         adj[to] = append(adj[to], from)
//         inDegrees[from]++
//     }

//     queue := make([]int, 0)
//     for i := range inDegrees {
//         if inDegrees[i] == 0 {
//             queue = append(queue, i)
//         }
//     }

//     for len(queue) > 0 {
//         course := queue[0]
//         queue = queue[1:]

//         result = append(result, course)

//         for _, nei := range adj[course] {
//             inDegrees[nei]--
//             if inDegrees[nei] == 0 {
//                 queue = append(queue, nei)
//             }
//         }
//     }

//     if len(result) == numCourses {
//         return result
//     }
//     return []int{}
// }
