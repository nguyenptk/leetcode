package medium

import (
	"testing"
)

func TestConstructorNumArray(t *testing.T) {
	cases := []struct {
		nums   []int
		sum1   []int
		want1  int
		upadte []int
		sum2   []int
		want2  int
	}{
		{
			[]int{1, 3, 5},
			[]int{0, 2},
			9,
			[]int{1, 2},
			[]int{0, 2},
			8,
		},
	}
	for _, c := range cases {
		obj := ConstructorNumArray(c.nums)
		got1 := obj.SumRange(c.sum1[0], c.sum1[1])
		if got1 != c.want1 {
			t.Errorf("ConstructorNumArray(%d) & SumRange(%d) == %d, want %d", c.nums, c.sum1, got1, c.want1)
		}

		obj.Update(c.upadte[0], c.upadte[1])
		got2 := obj.SumRange(c.sum2[0], c.sum2[1])
		if got1 != c.want1 {
			t.Errorf("ConstructorNumArray(%d) & SumRange(%d) == %d, want %d", c.nums, c.sum2, got2, c.want2)
		}
	}
}
