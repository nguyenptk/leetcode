// https://leetcode.com/problems/reverse-bits/
package easy

func reverseBits(num uint32) uint32 {
	result := uint32(0)

	for i := 0; i < 32; i++ {
		bit := (num >> 1) & 1
		result = result | bit<<(31-i)
	}

	return result
}
