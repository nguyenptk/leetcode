// https://leetcode.com/problems/minimum-genetic-mutation/
package medium

func MinMutation(start string, end string, bank []string) int {
	mapB := map[string]bool{}
	for _, v := range bank {
		mapB[v] = true
	}

	queue := []string{start}
	result := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			gene := queue[0]
			queue = queue[1:]
			if gene == end {
				return result
			}
			for j := 0; j < len(gene); j++ {
				for _, c := range "ACGT" {
					modStr := gene[:j] + string(c) + gene[j+1:]
					if mapB[modStr] {
						queue = append(queue, modStr)
						delete(mapB, modStr)
					}
				}
			}
		}
		result++
	}
	return -1
}
