// https://leetcode.com/problems/maximum-units-on-a-truck/
package easy

import "sort"

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	maxUnits := 0
	for _, v := range boxTypes {
		if truckSize >= v[0] {
			maxUnits += v[0] * v[1]
			truckSize -= v[0]
		} else {
			maxUnits += truckSize * v[1]
			break
		}
	}
	return maxUnits
}
