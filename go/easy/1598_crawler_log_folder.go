// https://leetcode.com/problems/crawler-log-folder/

package easy

func MinOperations(logs []string) int {
	count := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if count > 0 {
				count--
			}
		} else if logs[i] == "./" {
			continue
		} else {
			count++
		}
	}

	return count
}
