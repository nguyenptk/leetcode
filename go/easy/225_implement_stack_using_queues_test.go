package easy

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	cases := []struct {
		in   []int // push
		want map[string]string
	}{
		{
			[]int{1, 2},
			map[string]string{
				"top":   "2",     // equals to second push
				"pop":   "2",     // equals to second push
				"empty": "false", // not empty
			},
		},
		{
			[]int{1},
			map[string]string{
				"top":   "1",    // equals to first push
				"pop":   "1",    // equals to first push
				"empty": "true", // empty
			},
		},
		{
			[]int{10, -1, -5, 3, 7},
			map[string]string{
				"top":   "7",     // equals to last push
				"pop":   "7",     // equals to last push
				"empty": "false", // not empty
			},
		},
	}
	for _, c := range cases {
		myStack := Constructor()

		// push
		for _, v := range c.in {
			myStack.Push(v)
		}

		// call top function
		top := myStack.Top()
		if fmt.Sprint(top) != c.want["top"] {
			t.Errorf("Top == %s, want %s", fmt.Sprint(top), c.want["top"])
		}

		// call pop function
		pop := myStack.Pop()
		if fmt.Sprint(pop) != c.want["pop"] {
			t.Errorf("Pop == %s, want %s", fmt.Sprint(top), c.want["pop"])
		}

		// call empty function
		empty := myStack.Empty()
		if fmt.Sprint(empty) != c.want["empty"] {
			t.Errorf("Empty == %s, want %s", fmt.Sprint(empty), c.want["empty"])
		}
	}
}
