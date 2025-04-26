// https://leetcode.com/problems/nested-list-weight-sum/
package medium

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func DepthSum(nestedList []*NestedInteger) int {
	var dfs func(nestedList []*NestedInteger, depth int) int
	dfs = func(nestedList []*NestedInteger, depth int) int {
		if nestedList == nil {
			return 0
		}

		total := 0
		for _, v := range nestedList {
			if v.IsInteger() {
				total += v.GetInteger() * depth
			} else {
				total += dfs(v.GetList(), depth+1)
			}
		}

		return total
	}

	return dfs(nestedList, 1)
}
