// https://leetcode.com/problems/sentence-similarity-iii/
package medium

import "strings"

func AreSentencesSimilar(sentence1 string, sentence2 string) bool {
	dequeue1 := make([]string, 0)
	s1 := strings.Split(sentence1, " ")
	for i := 0; i < len(s1); i++ {
		dequeue1 = append(dequeue1, s1[i])
	}

	dequeue2 := make([]string, 0)
	s2 := strings.Split(sentence2, " ")
	for i := 0; i < len(s2); i++ {
		dequeue2 = append(dequeue2, s2[i])
	}

	for len(dequeue1) > 0 && len(dequeue2) > 0 &&
		dequeue1[0] == dequeue2[0] {
		dequeue1 = dequeue1[1:]
		dequeue2 = dequeue2[1:]
	}
	for len(dequeue1) > 0 && len(dequeue2) > 0 &&
		dequeue1[len(dequeue1)-1] == dequeue2[len(dequeue2)-1] {
		dequeue1 = dequeue1[:len(dequeue1)-1]
		dequeue2 = dequeue2[:len(dequeue2)-1]
	}

	if len(dequeue1) > 0 && len(dequeue2) > 0 {
		return false
	}
	return true
}
