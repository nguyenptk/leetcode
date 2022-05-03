package easy

import "testing"

func TestBackspaceCompare(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
	}
	for _, c := range cases {
		got := BackspaceCompare(c.in1, c.in2)
		if got != c.want {
			t.Errorf("BackspaceCompare(%q, %q) == %t, want %t", c.in1, c.in2, got, c.want)
		}
	}
}
