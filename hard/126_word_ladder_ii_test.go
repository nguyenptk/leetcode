package hard

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	cases := []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      [][]string
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}},
		},
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			[][]string{},
		},
	}
	for _, c := range cases {
		got := FindLadders(c.beginWord, c.endWord, c.wordList)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindLadders(%s, %s, %s) == %s, want %s", c.beginWord, c.endWord, c.wordList, got, c.want)
		}
	}
}
