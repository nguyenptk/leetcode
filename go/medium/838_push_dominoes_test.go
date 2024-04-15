package medium

import "testing"

func TestPushDominoes(t *testing.T) {
	cases := []struct {
		dominoes string
		want     string
	}{
		{"RR.L", "RR.L"},
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
	}
	for _, c := range cases {
		got := PushDominoes(c.dominoes)
		if got != c.want {
			t.Errorf("PushDominoes(%s) == %s, want %s", c.dominoes, got, c.want)
		}
	}
}
