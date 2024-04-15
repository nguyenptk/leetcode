package hard

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	cases := []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	}
	for _, c := range cases {
		got := FindSubstring(c.s, c.words)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindSubstring(%s, %s) == %d, want %d", c.s, c.words, got, c.want)
		}
	}
}
