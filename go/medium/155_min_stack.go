// https://leetcode.com/problems/min-stack
package medium

type MinStack struct {
	stack    []int
	minStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (f *MinStack) Push(val int) {
	f.stack = append(f.stack, val)
	if len(f.minStack) == 0 || val <= f.minStack[len(f.minStack)-1] {
		f.minStack = append(f.minStack, val)
	}
}

func (f *MinStack) Pop() {
	if f.stack[len(f.stack)-1] == f.minStack[len(f.minStack)-1] {
		f.minStack = f.minStack[:len(f.minStack)-1]
	}
	f.stack = f.stack[:len(f.stack)-1]
}

func (f *MinStack) Top() int {
	return f.stack[len(f.stack)-1]
}

func (f *MinStack) GetMin() int {
	if len(f.minStack) > 0 {
		return f.minStack[len(f.minStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
