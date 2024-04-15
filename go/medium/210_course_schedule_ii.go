// https://leetcode.com/problems/course-schedule
package medium

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)

	for _, p := range prerequisites {
		from := p[1]
		to := p[0]
		graph[from] = append(graph[from], to)
	}

	visit := make(map[int]bool)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if visit[course] {
			return false
		}
		if len(graph[course]) == 0 {
			return true
		}

		visit[course] = true

		for _, p := range graph[course] {
			if !dfs(p) {
				return false
			}
		}
		delete(visit, course)
		graph[course] = []int{}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
