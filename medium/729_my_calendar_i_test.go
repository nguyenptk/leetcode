package medium

import (
	"reflect"
	"testing"
)

func TestConstructorMyCalendar(t *testing.T) {
	cases := []struct {
		books [][]int
		want  []bool
	}{
		{
			[][]int{{10, 20}, {15, 25}, {20, 30}},
			[]bool{true, false, true},
		},
		{
			[][]int{{47, 50}, {33, 41}, {39, 45}, {33, 42}, {25, 32}, {26, 35}, {19, 25}, {3, 8}, {8, 13}, {18, 27}},
			[]bool{true, true, false, false, true, false, true, true, true, false},
		},
	}
	for _, c := range cases {
		obj := ConstructorMyCalendar()
		var got []bool
		for _, book := range c.books {
			got = append(got, obj.Book(book[0], book[1]))
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ConstructorMyCalendar() == %t, want %t", got, c.want)
		}
	}
}
