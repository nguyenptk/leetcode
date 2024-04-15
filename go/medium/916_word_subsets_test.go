package medium

import (
	"reflect"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	cases := []struct {
		words1 []string
		words2 []string
		want   []string
	}{
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}, []string{"apple", "google", "leetcode"}},
	}
	for _, c := range cases {
		got := WordSubsets(c.words1, c.words2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("WordSubsets(%s, %s) == %s, want %s", c.words1, c.words2, got, c.want)
		}
	}
}
