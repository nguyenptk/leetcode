package medium

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		s    string
		k    int
		want string
	}{
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
		{"nnwssswwnvbnnnbbqhhbbbhmmmlllm", 3, "vqm"},
		{"yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy", 4, "ybth"},
	}
	for _, c := range cases {
		got := RemoveDuplicates(c.s, c.k)
		if got != c.want {
			t.Errorf("RemoveDuplicates(%s, %d) == %s, want %s", c.s, c.k, got, c.want)
		}
	}
}
