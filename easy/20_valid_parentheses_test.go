package easy

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"{}[]()(())", true},
		{"(]", false},
		{"{                }", true},
	}
	for _, c := range cases {
		got := IsValid(c.in)
		if got != c.want {
			t.Errorf("IsValid(%q) == %t, want %t", c.in, got, c.want)
		}
	}

}
