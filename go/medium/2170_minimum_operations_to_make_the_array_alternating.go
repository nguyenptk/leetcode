// https://leetcode.com/problems/minimum-operations-to-make-the-array-alternating/
package medium

func MinimumOperations(nums []int) int {
	evenFreq := make(map[int]int)
	oddFreq := make(map[int]int)
	evenCount, oddCount := 0, 0

	for i, num := range nums {
		if i%2 == 0 {
			evenFreq[num]++
			evenCount++
		} else {
			oddFreq[num]++
			oddCount++
		}
	}

	even1Val, even1Cnt, even2Cnt := 0, 0, 0
	for val, cnt := range evenFreq {
		if cnt > even1Cnt {
			even2Cnt = even1Cnt
			even1Cnt = cnt
			even1Val = val
		} else if cnt > even2Cnt {
			even2Cnt = cnt
		}
	}

	odd1Val, odd1Cnt, odd2Cnt := 0, 0, 0
	for val, cnt := range oddFreq {
		if cnt > odd1Cnt {
			odd2Cnt = odd1Cnt
			odd1Cnt = cnt
			odd1Val = val
		} else if cnt > odd2Cnt {
			odd2Cnt = cnt
		}
	}

	if even1Val != odd1Val {
		return (evenCount - even1Cnt) + (oddCount - odd1Cnt)
	}

	opt1 := (evenCount - even1Cnt) + (oddCount - odd2Cnt)
	opt2 := (evenCount - even2Cnt) + (oddCount - odd1Cnt)

	return min(opt1, opt2)
}

// Time complexity: O(n)
// Space complexity: O(n)
