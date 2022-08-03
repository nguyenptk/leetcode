// https://leetcode.com/problems/my-calendar-i/
package medium

import "math"

type MyCalendar struct {
	start int
	end   int
	next  *MyCalendar
}

func ConstructorMyCalendar() MyCalendar {
	return MyCalendar{start: -1, end: -1, next: &MyCalendar{start: math.MaxInt, end: math.MaxInt, next: nil}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	curr := this
	last := curr
	for start >= curr.end {
		last = curr
		curr = curr.next
	}
	if curr.start < end {
		return false
	}
	last.next = &MyCalendar{start: start, end: end, next: curr}
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
