// https://leetcode.com/problems/moving-average-from-data-stream/
package easy

type MovingAverage struct {
	queue []int
	size  int
	sum   int
}

func ConstructorMovingAverage(size int) MovingAverage {
	return MovingAverage{
		queue: make([]int, 0),
		size:  size,
		sum:   0,
	}
}

func (f *MovingAverage) Next(val int) float64 {
	f.queue = append(f.queue, val)
	f.sum += val

	for len(f.queue) > f.size {
		f.sum -= f.queue[0]
		f.queue = f.queue[1:]
	}

	return float64(f.sum) / float64(len(f.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
