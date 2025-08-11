// https://leetcode.com/problems/design-excel-sum-formula/
package hard

type Cell struct {
	val      int
	children []*Cell
	parents  []*Cell
	cached   *int
}

func (c *Cell) getValue() int {
	if c.cached != nil {
		return *c.cached
	}
	if c.children == nil {
		c.cached = new(int)
		*c.cached = c.val
		return *c.cached
	}

	sum := 0
	for _, child := range c.children {
		sum += child.getValue()
	}

	c.cached = new(int)
	*c.cached = sum
	return sum
}

type Excel struct {
	matrix [][]*Cell
}

func ConstructorExcel(height int, width byte) Excel {
	matrix := make([][]*Cell, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([]*Cell, width-'A'+1)
		for j := range matrix[i] {
			matrix[i][j] = &Cell{}
		}
	}
	return Excel{matrix: matrix}
}

func (f *Excel) Set(row int, column byte, val int) {
	cell := f.matrix[row-1][column-'A']
	cell.val = val
	f.clearChildren(cell)
	f.invalidate(cell)
}

func (f *Excel) Get(row int, column byte) int {
	return f.matrix[row-1][column-'A'].getValue()
}

func (f *Excel) Sum(row int, column byte, numbers []string) int {
	cell := f.matrix[row-1][column-'A']

	f.clearChildren(cell) // Remove old dependency links

	for _, s := range numbers {
		if idx := indexOf(s, ":"); idx != -1 {
			c1, r1 := s[0], toInt(s[1:idx])
			c2, r2 := s[idx+1], toInt(s[idx+2:])
			for i := r1 - 1; i < r2; i++ {
				for j := c1 - 'A'; j <= c2-'A'; j++ {
					child := f.matrix[i][j]
					cell.children = append(cell.children, child)
					child.parents = append(child.parents, cell)
				}
			}
		} else {
			c, r := s[0], toInt(s[1:])
			child := f.matrix[r-1][c-'A']
			cell.children = append(cell.children, child)
			child.parents = append(child.parents, cell)
		}
	}

	f.invalidate(cell)
	return cell.getValue()
}

func (f *Excel) invalidate(cell *Cell) {
	if cell.cached == nil {
		return
	}
	cell.cached = nil
	for _, parent := range cell.parents {
		f.invalidate(parent)
	}
}

func (f *Excel) clearChildren(cell *Cell) {
	// When re-setting or re-summing a node, remove it from its children's parent lists
	if cell.children != nil {
		for _, child := range cell.children {
			newParents := make([]*Cell, 0)
			for _, p := range child.parents {
				if p != cell {
					newParents = append(newParents, p)
				}
			}
			child.parents = newParents
		}
		cell.children = nil
	}
}

func indexOf(s, sep string) int {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == sep {
			return i
		}
	}
	return -1
}

func toInt(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num
}
