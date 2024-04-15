package easy

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, c := range cases {
		got := IsAnagram(c.s, c.t)
		if got != c.want {
			t.Errorf("IsAnagram(%s, %s) == %t, want %t", c.s, c.t, got, c.want)
		}
	}
}
