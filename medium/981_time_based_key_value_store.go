// https://leetcode.com/problems/time-based-key-value-store/
package medium

type TimeMap struct {
	timeMap map[string][]data
}

type data struct {
	timestamp int
	value     string
}

func ConstructorTimeMap() TimeMap {
	return TimeMap{
		map[string][]data{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMap[key] = append(this.timeMap[key], data{timestamp: timestamp,
		value: value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var list []data
	var ok bool
	if list, ok = this.timeMap[key]; !ok {
		return ""
	}

	l := 0
	r := len(list)

	for l < r {
		m := (l + r) / 2
		if list[m].timestamp > timestamp {
			r = m
		} else {
			l = m + 1
		}
	}

	if l == 0 {
		return ""
	}
	return list[l-1].value

}
