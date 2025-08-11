// https://leetcode.com/problems/maximize-ysum-by-picking-a-triplet-of-distinct-xvalues/
package medium

import "sort"

func MaxSumDistinctTriplet(x []int, y []int) int {
	topY := make(map[int]int)
	for i := range x {
		if curr, ok := topY[x[i]]; !ok || y[i] > curr {
			topY[x[i]] = y[i]
		}
	}

	vals := make([]int, 0)
	for _, v := range topY {
		vals = append(vals, v)
	}
	if len(vals) < 3 {
		return -1
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	return vals[0] + vals[1] + vals[2]
}

// Time complexity: O(n + k log k)
// Space complexity: O(n)

// func MaxSumDistinctTriplet(x []int, y []int) int {
//     n := len(x)

//     var a, b, c int
//     var a1, b1, c1 int
//     for i := 0; i < n; i++ {
//         if x[i] == a1 {
//             a = max(a, y[i])
//         } else if x[i] == b1 {
//             b = max(b, y[i])
//         } else if x[i] == c1 {
//             c = max(c, y[i])
//         } else if y[i] > a || y[i] > b || y[i] > c {
//             mi := min(a, min(b, c))
//             if mi == a {
//                 a = y[i]
//                 a1 = x[i]
//             } else if mi == b {
//                 b = y[i]
//                 b1 = x[i]
//             } else {
//                 c = y[i]
//                 c1 = x[i]
//             }
//         }
//     }

//     if a == 0 || b == 0 || c == 0 {
//         return -1
//     }

//     return a+b+c
// }

// Time complexity: O(n)
// Space complexity: O(1)
