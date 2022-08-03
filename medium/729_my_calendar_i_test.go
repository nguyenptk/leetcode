package medium

import (
	"reflect"
	"testing"
)

func TestConstructorMyCalendar(t *testing.T) {
	cases := []struct {
		book1 []int
		book2 []int
		book3 []int
		want  []bool
	}{
		{
			[]int{10, 20},
			[]int{15, 25},
			[]int{20, 30},
			[]bool{true, false, true},
		},
	}
	for _, c := range cases {
		obj := ConstructorMyCalendar()
		var got []bool
		got = append(got, obj.Book(c.book1[0], c.book1[1]))
		got = append(got, obj.Book(c.book2[0], c.book2[1]))
		got = append(got, obj.Book(c.book3[0], c.book3[1]))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ConstructorMyCalendar() == %t, want %t", got, c.want)
		}
	}
}
