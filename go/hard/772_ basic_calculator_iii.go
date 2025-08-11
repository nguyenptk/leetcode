// https://leetcode.com/problems/basic-calculator-iii/
package hard

import "unicode"

func CalculateIII(s string) int {
	s += "0"

	i := 0
	var rec func() int
	rec = func() int {
		operator, result, curr, last := byte('+'), 0, 0, 0
		for i < len(s) {
			ch := s[i]
			i++

			if unicode.IsDigit(rune(ch)) {
				curr = curr*10 + int(ch-'0')
			}
			if ch == '(' {
				curr = rec()
			}
			if !unicode.IsDigit(rune(ch)) || i == len(s)-1 {
				switch operator {
				case '+':
					result += last
					last = curr
				case '-':
					result += last
					last = -curr
				case '*':
					last *= curr
				case '/':
					last /= curr
				}

				if ch == ')' {
					break
				}
				curr = 0
				operator = ch
			}
		}
		return result + last
	}

	return rec()
}

// Time Complexity:	O(n) - Even for nested parentheses, the recursive call advances the shared index i, so characters aren't reprocessed.
// Worst case: O(n) - All parentheses are nested

// Stack approach
// func CalculateIII(s string) int {
//     // three parallel stacks for (,)
//     var stackRes, stackLast []int
//     var stackOp        []byte

//     result, last, curr := 0, 0, 0
//     operator := byte('+')

//     for i := 0; i < len(s); i++ {
//         ch := s[i]

//         // 1) skip spaces
//         if ch == ' ' {
//             continue
//         }
//         // 2) accumulate multi‑digit numbers
//         if ch >= '0' && ch <= '9' {
//             curr = curr*10 + int(ch-'0')
//             continue
//         }
//         // 3) on '(', push all state
//         if ch == '(' {
//             stackRes  = append(stackRes,  result)
//             stackLast = append(stackLast, last)
//             stackOp   = append(stackOp,   operator)
//             // reset for sub‑expr
//             result, last, curr = 0, 0, 0
//             operator = '+'
//             continue
//         }

//         // 4) on any operator or ')', first apply the **previous** operator
//         switch operator {
//         case '+':
//             result += last
//             last = curr
//         case '-':
//             result += last
//             last = -curr
//         case '*':
//             last *= curr
//         case '/':
//             last /= curr
//         }
//         curr = 0

//         // 5) if this is ')', pop state and treat the sub‑result as curr
//         if ch == ')' {
//             subtotal := result + last
//             // pop in reverse order
//             L := len(stackRes) - 1
//             result   = stackRes[L]
//             last     = stackLast[L]
//             operator = stackOp[L]
//             stackRes  = stackRes[:L]
//             stackLast = stackLast[:L]
//             stackOp   = stackOp[:L]
//             // carry the computed subtotal back as curr for the outer frame
//             curr = subtotal
//             continue
//         }

//         // 6) otherwise it’s one of '+','-','*','/'
//         operator = ch
//     }

//     // 7) apply the very last pending operation
//     switch operator {
//     case '+':
//         result += last
//         last = curr
//     case '-':
//         result += last
//         last = -curr
//     case '*':
//         last *= curr
//     case '/':
//         last /= curr
//     }

//     return result + last
// }
