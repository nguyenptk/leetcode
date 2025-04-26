// https://leetcode.com/problems/employee-free-time/
package hard

import "sort"

type Interval struct {
	Start int
	End   int
}

func EmployeeFreeTime(schedule [][]*Interval) []*Interval {
	intervals := make([]*Interval, 0)
	for _, sch := range schedule {
		intervals = append(intervals, sch...)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := make([]*Interval, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if last.End < intervals[i].Start {
			merged = append(merged, intervals[i])
		} else {
			last.End = max(last.End, intervals[i].End)
		}
	}

	result := make([]*Interval, 0)
	for i := 0; i < len(merged)-1; i++ {
		interval := &Interval{
			Start: merged[i].End,
			End:   merged[i+1].Start,
		}
		result = append(result, interval)
	}

	return result
}
