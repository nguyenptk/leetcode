// https://leetcode.com/problems/satisfiability-of-equality-equations/
package medium

type UnionFind struct {
	ID []int
}

func (f *UnionFind) UnionFind(n int) {
	f.ID = make([]int, n)
	for i := 0; i < n; i++ {
		f.ID[i] = i
	}
}

func (f *UnionFind) union(u, v int) {
	f.ID[f.find(u)] = f.find(v)
}

func (f *UnionFind) find(u int) int {
	if f.ID[u] == u {
		return u
	}
	f.ID[u] = f.find(f.ID[u])
	return f.ID[u]
}

func EquationsPossible(equations []string) bool {
	f := &UnionFind{}
	f.UnionFind(26)

	for _, e := range equations {
		if e[1] == '=' {
			x := e[0] - 'a'
			y := e[3] - 'a'
			f.union(int(x), int(y))
		}
	}

	for _, e := range equations {
		if e[1] == '!' {
			x := e[0] - 'a'
			y := e[3] - 'a'
			if f.find(int(x)) == f.find(int(y)) {
				return false
			}
		}
	}
	return true
}
