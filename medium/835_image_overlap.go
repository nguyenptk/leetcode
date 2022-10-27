// https://leetcode.com/problems/image-overlap/
package medium

func LargestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	magic := 100
	result := 0
	ones1 := [][]int{}
	ones2 := [][]int{}
	mapData := map[int]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] == 1 {
				ones1 = append(ones1, [][]int{{i, j}}...)
			}
			if img2[i][j] == 1 {
				ones2 = append(ones2, [][]int{{i, j}}...)
			}
		}
	}

	for _, v1 := range ones1 {
		for _, v2 := range ones2 {
			key := (v1[0]-v2[0])*magic + v1[1] - v2[1]
			if _, ok := mapData[key]; ok {
				mapData[key] += 1
			} else {
				mapData[key] = 1
			}
		}
	}

	for _, v := range mapData {
		if result < v {
			result = v
		}
	}

	return result
}
