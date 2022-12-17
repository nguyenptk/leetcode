package medium

import "testing"

func TestEvalRPN(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, c := range cases {
		got := EvalRPN(c.in)
		if got != c.want {
			t.Errorf("EvalRPN(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
