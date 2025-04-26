// https://leetcode.com/problems/course-schedule
package medium

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for _, v := range prerequisites {
		from := v[0]
		to := v[1]
		graph[from] = append(graph[from], to)
	}

	for i := 0; i < numCourses; i++ {
		if !dfsCanFinish(graph, i, make(map[int]bool)) {
			return false
		}
	}

	return true

}

func dfsCanFinish(graph map[int][]int, course int, visited map[int]bool) bool {
	if visited[course] {
		return false
	}
	if len(graph[course]) == 0 {
		return true
	}

	visited[course] = true

	for _, p := range graph[course] {
		if !dfsCanFinish(graph, p, visited) {
			return false
		}
	}

	delete(visited, course)
	graph[course] = []int{}

	return true
}

// Topology Sort approach
// func CanFinish(numCourses int, prerequisites [][]int) bool {
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

//     coursesTaken := 0
//     for len(queue) > 0 {
//         course := queue[0]
//         queue = queue[1:]
//         coursesTaken += 1
//         for _, nei := range adj[course] {
//             inDegrees[nei] -= 1
//             if inDegrees[nei] == 0 {
//                 queue = append(queue, nei)
//             }
//         }
//     }

//     return coursesTaken == numCourses
// }
