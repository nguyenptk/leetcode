package medium

import "testing"

func TestMinMutation(t *testing.T) {
	cases := []struct {
		start string
		end   string
		bank  []string
		want  int
	}{
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
	}
	for _, c := range cases {
		got := MinMutation(c.start, c.end, c.bank)
		if got != c.want {
			t.Errorf("MinMutation(%s, %s, %s) == %d, want %d", c.start, c.end, c.bank, got, c.want)
		}
	}
}
