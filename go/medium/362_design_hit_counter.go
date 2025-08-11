// https://leetcode.com/problems/design-hit-counter/
package medium

type Hit struct {
	time  int
	count int
}

type HitCounter struct {
	total int
	hits  []*Hit
}

func ConstructorHitCounter() HitCounter {
	return HitCounter{
		total: 0,
		hits:  make([]*Hit, 0),
	}
}

func (f *HitCounter) Hit(timestamp int) {
	if len(f.hits) == 0 || f.hits[len(f.hits)-1].time != timestamp {
		f.hits = append(f.hits, &Hit{time: timestamp, count: 1})
	} else {
		prevCount := f.hits[len(f.hits)-1].count
		f.hits = f.hits[:len(f.hits)-1]
		f.hits = append(f.hits, &Hit{time: timestamp, count: prevCount + 1})
	}
	f.total++
}

func (f *HitCounter) GetHits(timestamp int) int {
	for len(f.hits) > 0 {
		diff := timestamp - f.hits[0].time
		if diff >= 300 {
			f.total -= f.hits[0].count
			f.hits = f.hits[1:]
		} else {
			break
		}
	}
	return f.total
}

// Time Complexity:
// hit - O(1).
// getHits - If there are a total of n pairs present in the deque, worst case time complexity can be O(n). However, by clubbing all the timestamps with same value together, for the ith timestamp with k repetitions, the time complexity is O(1) as here, instead of removing all those k repetitions, we only remove a single entry from the deque.

// Space complexity:
// If there are a total of N elements that we encountered throughout, the space complexity is O(N) (similar to Approach 1). However, in the case of repetitions, the space required for storing those k values O(1).

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
