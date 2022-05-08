package medium

import (
	"reflect"
	"strconv"
	"testing"
)

func TestConstructor(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{
			"[[1,1],2,[1,1]]",
			[]int{1, 1, 2, 1, 1},
		},
		{
			"[1,[4,[6]]]",
			[]int{1, 4, 6},
		},
	}
	for _, c := range cases {
		nestedInteger := deserialize(c.in)
		arr := []*NestedInteger{nestedInteger}
		iterator := Constructor(arr)
		got := []int{}
		for iterator.HasNext() {
			got = append(got, iterator.Next())
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Constructor(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}

// Parse string to NestedInteger
func deserialize(s string) *NestedInteger {
	stack, cur := []*NestedInteger{}, &NestedInteger{}
	for i := 0; i < len(s); {
		switch {
		case isDigital(s[i]) || s[i] == '-':
			j := 0
			for j = i + 1; j < len(s) && isDigital(s[j]); j++ {
			}
			num, _ := strconv.Atoi(s[i:j])
			next := &NestedInteger{}
			next.SetInteger(num)
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			} else {
				cur = next
			}
			i = j
		case s[i] == '[':
			next := &NestedInteger{}
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			}
			stack = append(stack, next)
			i++
		case s[i] == ']':
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
		case s[i] == ',':
			i++
		}
	}
	return cur
}

// Check byte is integer
func isDigital(v byte) bool {
	if v >= '0' && v <= '9' {
		return true
	}
	return false
}
