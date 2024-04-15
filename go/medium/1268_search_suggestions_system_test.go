package medium

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	cases := []struct {
		products   []string
		searchWord string
		want       [][]string
	}{
		// Test case 1
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		// Test case 2
		{
			[]string{"havana"},
			"havana",
			[][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			},
		},
		// Test case 3
		{
			[]string{"bags", "baggage", "banner", "box", "cloths"},
			"bags",
			[][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
	}
	for _, c := range cases {
		got := SuggestedProducts(c.products, c.searchWord)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SuggestedProducts(%s, %s) == %s, want %s", c.products, c.searchWord, got, c.want)
		}
	}
}
