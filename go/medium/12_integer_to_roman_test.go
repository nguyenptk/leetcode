package medium

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for _, c := range cases {
		got := IntToRoman(c.in)
		if got != c.want {
			t.Errorf("IntToRoman(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}
