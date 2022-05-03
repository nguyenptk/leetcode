package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, c := range cases {
		got := LongestCommonPrefix(c.in)
		if got != c.want {
			t.Errorf("LongestCommonPrefix(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
