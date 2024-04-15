// https://leetcode.com/problems/reordered-power-of-2/
package medium

func ReorderedPowerOf2(n int) bool {
	d := getNumOfDigits(n)
	a := make([]int, 10)
	for v := range a {
		a[v] = 0
	}
	getDigitsFreq(n, a)
	n2 := 1
	var d1 int
	for i := 0; i < 30; i++ {
		d1 = getNumOfDigits(n2)
		if d == d1 {
			a1 := make([]int, 10)
			for v := range a1 {
				a1[v] = 0
			}
			getDigitsFreq(n2, a1)
			var j int
			for j = 0; j < 10; j++ {
				if a[j] != a1[j] {
					break
				}
			}
			if j == 10 {
				return true
			}
		} else if d < d1 {
			break
		}
		n2 = n2 << 1
	}
	return false
}

func getDigitsFreq(n int, a []int) {
	for n > 0 {
		a[n%10]++
		n /= 10
	}
}

func getNumOfDigits(n int) int {
	i := 0
	for n > 0 {
		i++
		n /= 10
	}
	return i
}
