// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
package hard

import "sort"

type Seed struct {
	p int
	g int
}

func EarliestFullBloom(plantTime []int, growTime []int) int {
	result := 0
	time := 0
	seeds := []Seed{}

	for i := 0; i < len(plantTime); i++ {
		seeds = append(seeds, Seed{p: plantTime[i], g: growTime[i]})
	}
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].g > seeds[j].g
	})

	for _, seed := range seeds {
		time += seed.p
		if result < time+seed.g {
			result = time + seed.g
		}
	}

	return result
}
