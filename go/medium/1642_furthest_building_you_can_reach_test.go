package medium

import "testing"

func TestFurthestBuilding(t *testing.T) {
	cases := []struct {
		heights []int
		bricks  int
		ladders int
		want    int
	}{
		{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{[]int{14, 3, 19, 3}, 17, 0, 3},
	}
	for _, c := range cases {
		got := FurthestBuilding(c.heights, c.bricks, c.ladders)
		if got != c.want {
			t.Errorf("FurthestBuilding(%d, %d, %d) == %d, want %d", c.heights, c.bricks, c.ladders, got, c.want)
		}
	}
}
