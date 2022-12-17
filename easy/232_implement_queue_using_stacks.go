// https://leetcode.com/problems/implement-queue-using-stacks
package easy

type MyQueue struct {
	nodes []*Node
	len   int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.nodes = append(this.nodes,
		&Node{
			val: x,
		})
	this.len++
}

func (this *MyQueue) Pop() int {
	if !this.Empty() {
		this.len--
		val := this.nodes[0].val    // get first node
		this.nodes = this.nodes[:1] // pop node
		return val
	}
	return -1
}

func (this *MyQueue) Peek() int {
	return this.nodes[0].val
}

func (this *MyQueue) Empty() bool {
	if this.len > 0 && len(this.nodes) > 0 {
		return false
	}
	return true
}
