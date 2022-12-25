package easy

import (
	"reflect"
	"testing"
)

func TestAnswerQueries(t *testing.T) {
	cases := []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
		{[]int{2, 3, 4, 5}, []int{1}, []int{0}},
	}
	for _, c := range cases {
		got := AnswerQueries(c.nums, c.queries)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AnswerQueries(%d, %d) == %d, want %d", c.nums, c.queries, got, c.want)
		}
	}
}
