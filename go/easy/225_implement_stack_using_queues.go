// https://leetcode.com/problems/implement-stack-using-queues/

package easy

type MyStack struct {
	nodes []*Node
	len   int
}

type Node struct {
	val int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.nodes = append(this.nodes,
		&Node{
			val: x,
		})
	this.len++
}

func (this *MyStack) Pop() int {
	this.len--
	val := this.nodes[this.len].val    // get last node
	this.nodes = this.nodes[:this.len] // pop node
	return val
}

func (this *MyStack) Top() int {
	return this.nodes[this.len-1].val
}

func (this *MyStack) Empty() bool {
	if this.len > 0 {
		return false
	}
	return true
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
