// https://leetcode.com/problems/can-place-flowers/
package easy

func CanPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			emptyL := (i == 0) || (flowerbed[i-1] == 0)
			emptyR := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			if emptyL && emptyR {
				flowerbed[i] = 1
				n--
				if n == 0 {
					return true
				}
			}
		}
	}

	return n <= 0
}
