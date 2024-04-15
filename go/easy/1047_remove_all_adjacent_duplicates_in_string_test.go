package easy

import "testing"

func TestRemoveAllAdjacentDuplicates(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
	}
	for _, c := range cases {
		got := RemoveAllAdjacentDuplicates(c.s)
		if got != c.want {
			t.Errorf("RemoveAllAdjacentDuplicates(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
