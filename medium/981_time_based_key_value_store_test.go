package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructorTimeMap(t *testing.T) {
	timeMap := ConstructorTimeMap()
	timeMap.Set("foo", "bar", 1)                   // store the key "foo" and value "bar" along with timestamp = 1.
	assert.Equal(t, timeMap.Get("foo", 1), "bar")  // return "bar"
	assert.Equal(t, timeMap.Get("foo", 1), "bar")  // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	timeMap.Set("foo", "bar2", 4)                  // store the key "foo" and value "bar2" along with timestamp = 4.
	assert.Equal(t, timeMap.Get("foo", 4), "bar2") // return "bar2"
	assert.Equal(t, timeMap.Get("foo", 5), "bar2") // return "bar2"
}
