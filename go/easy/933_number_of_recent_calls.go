// https://leetcode.com/problems/number-of-recent-calls/
package easy

type RecentCounter struct {
	stack []int
}

func ConstructorRecentCounter() RecentCounter {
	return RecentCounter{
		stack: make([]int, 0),
	}
}

func (f *RecentCounter) Ping(t int) int {
	f.stack = append(f.stack, t)
	for f.stack[0] < t-3000 {
		f.stack = f.stack[1:]
	}
	return len(f.stack)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
