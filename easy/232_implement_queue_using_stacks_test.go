package easy

import (
	"fmt"
	"testing"
)

func TestConstructorMyQueue(t *testing.T) {
	cases := []struct {
		in   []int // push
		want map[string]string
	}{
		{
			[]int{1, 2},
			map[string]string{
				"peek":  "1",     // equals to second push
				"pop":   "1",     // equals to second push
				"empty": "false", // not empty
			},
		},
		{
			[]int{1},
			map[string]string{
				"peek":  "1",    // equals to first push
				"pop":   "1",    // equals to first push
				"empty": "true", // empty
			},
		},
		{
			[]int{10, -1, -5, 3, 7},
			map[string]string{
				"peek":  "10",    // equals to last push
				"pop":   "10",    // equals to last push
				"empty": "false", // not empty
			},
		},
	}
	for _, c := range cases {
		myQueue := ConstructorMyQueue()

		// push
		for _, v := range c.in {
			myQueue.Push(v)
		}

		// call peek function
		peek := myQueue.Peek()
		if fmt.Sprint(peek) != c.want["peek"] {
			t.Errorf("Top == %s, want %s", fmt.Sprint(peek), c.want["peek"])
		}

		// call pop function
		pop := myQueue.Pop()
		if fmt.Sprint(pop) != c.want["pop"] {
			t.Errorf("Pop == %s, want %s", fmt.Sprint(peek), c.want["pop"])
		}

		// call empty function
		empty := myQueue.Empty()
		if fmt.Sprint(empty) != c.want["empty"] {
			t.Errorf("Empty == %s, want %s", fmt.Sprint(empty), c.want["empty"])
		}
	}
}
