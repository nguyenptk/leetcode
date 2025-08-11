// https://leetcode.com/problems/employee-importance/
package medium

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func GetImportance(employees []*Employee, id int) int {
	m := map[int]*Employee{}
	for _, e := range employees {
		m[e.Id] = e
	}
	var dfs func(eid int) int
	dfs = func(eid int) int {
		e := m[eid]
		result := e.Importance
		for _, subid := range e.Subordinates {
			result += dfs(subid)
		}
		return result
	}
	return dfs(id)
}
