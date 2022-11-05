package hard

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	cases := []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{
			[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"oath", "eat"},
		},
		{
			[][]byte{{'a', 'b'}, {'c', 'd'}},
			[]string{"abcb"},
			[]string{},
		},
	}
	for _, c := range cases {
		got := FindWords(c.board, c.words)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindWords(%s, %s) == %s, want %s", c.board, c.words, got, c.want)
		}
	}
}
