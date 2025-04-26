package hard

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, c := range cases {
		got := CalculateI(c.s)
		if got != c.want {
			t.Errorf("Calculate(%s) == %d, want %d", c.s, got, c.want)
		}
	}
}
