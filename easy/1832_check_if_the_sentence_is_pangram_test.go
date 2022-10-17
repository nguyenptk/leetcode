package easy

import (
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	cases := []struct {
		sentence string
		want     bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
	}
	for _, c := range cases {
		got := CheckIfPangram(c.sentence)
		if got != c.want {
			t.Errorf("CheckIfPangram(%s) == %t, want %t", c.sentence, got, c.want)
		}
	}
}
