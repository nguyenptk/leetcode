// https://leetcode.com/problems/find-unique-binary-string/
package medium

func FindDifferentBinaryString(nums []string) string {
	str := ""
	for i := range nums {
		curr := nums[i][i]
		if curr == '0' {
			str += "1"
		} else {
			str += "0"
		}

	}

	return str
}
