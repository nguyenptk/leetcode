// https://leetcode.com/problems/encode-and-decode-strings/
package medium

import "strings"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	return strings.Join(strs, "\n")
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return strings.Split(strs, "\n")
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
