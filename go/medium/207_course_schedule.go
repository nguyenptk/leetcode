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
