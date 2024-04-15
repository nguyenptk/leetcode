package easy

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, c := range cases {
		got := RomanToInt(c.in)
		if got != c.want {
			t.Errorf("RomanToInt(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
