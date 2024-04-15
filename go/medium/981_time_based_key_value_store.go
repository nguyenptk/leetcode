// https://leetcode.com/problems/time-based-key-value-store/
package medium

type TimeMap struct {
	mapData map[string][]data
}

type data struct {
	value     string
	timestamp int
}

func ConstructorTimeMap() TimeMap {
	return TimeMap{
		map[string][]data{},
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	data := data{
		value:     value,
		timestamp: timestamp,
	}
	t.mapData[key] = append(t.mapData[key], data)
}

func (t *TimeMap) Get(key string, timestamp int) string {
	list := t.mapData[key]

	// Binary search
	l := 0
	r := len(list) - 1
	result := ""
	for l <= r {
		m := (l + r) / 2
		if list[m].timestamp <= timestamp {
			l = m + 1
			result = list[m].value
		} else {
			r = m - 1
		}
	}

	return result
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
