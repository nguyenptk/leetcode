// https://leetcode.com/problems/flatten-nested-list-iterator/
package medium

// This is the interface that allows for creating nested lists.
type NestedInteger struct {
	Number int
	List   []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	if this.List == nil {
		return true
	}
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.Number
}

// Set this NestedInteger to hold a single integer.
func (this *NestedInteger) SetInteger(value int) {
	this.Number = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	this.List = append(this.List, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.List
}

type NestedIterator struct {
	// init stack is a queue
	stack [][]*NestedInteger
}

func Constructor(stack []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: [][]*NestedInteger{stack},
	}
}

func (this *NestedIterator) Next() int {
	// Get the first element of Top in stack list
	queue := this.stack[len(this.stack)-1]
	// Get the first value of queue and return it
	val := queue[0].GetInteger()
	// Pop the first value to the head of queue
	this.stack[len(this.stack)-1] = queue[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	// Check next of stack
	for len(this.stack) > 0 {
		// Get the first element of Top in stack list
		queue := this.stack[len(this.stack)-1]
		if len(queue) == 0 {
			// Pop the stack
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		// Assign the first element of queue to nest
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		// if the first element is a list and then merge pop-up queue into the stack
		this.stack[len(this.stack)-1] = queue[1:]
		this.stack = append(this.stack, nest.GetList())
	}
	return false
}
