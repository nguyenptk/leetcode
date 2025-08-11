package medium

func MaxValueInABoundedArry(n int, index int, value int) int {
	l, r := 1, value
	for l < r {
		m := (l + r + 1) / 2
		if getSumInABoundedArry(n, index, m) <= value {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func getSumInABoundedArry(n int, index int, value int) int {
	count := 0
	if value > index {
		count += (value + value - index) * (index + 1) / 2
	} else {
		count += (value+1)*value/2 + index - value + 1
	}
	if value >= n-index {
		count += (value + value - n + 1 + index) * (n - index) / 2
	} else {
		count += (value+1)*value/2 + n - index - value
	}

	return count - value
}
