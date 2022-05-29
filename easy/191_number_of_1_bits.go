// https://leetcode.com/problems/number-of-1-bits/
package easy

import "math/bits"

func HammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
