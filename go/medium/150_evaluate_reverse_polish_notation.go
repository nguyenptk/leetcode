// https://leetcode.com/problems/evaluate-reverse-polish-notation
package medium

import "strconv"

func EvalRPN(tokens []string) int {
	result := []int{} // stack
	for _, v := range tokens {
		switch {
		case v == "+":
			s1 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			s2 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			result = append(result, s1+s2)
			break
		case v == "-":
			s1 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			s2 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			result = append(result, s2-s1)
			break
		case v == "*":
			s1 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			s2 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			result = append(result, s1*s2)
			break
		case v == "/":
			s1 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			s2 := result[len(result)-1]
			result = result[:len(result)-1] // POP
			result = append(result, s2/s1)
			break
		default:
			n, _ := strconv.Atoi(v)
			result = append(result, n)
		}
	}
	return result[len(result)-1]
}
