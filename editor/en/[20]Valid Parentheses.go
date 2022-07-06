//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']
//', determine if the input string is valid.
//
// An input string is valid if:
//
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
//
//
// Example 1:
//
//
//Input: s = "()"
//Output: true
//
//
// Example 2:
//
//
//Input: s = "()[]{}"
//Output: true
//
//
// Example 3:
//
//
//Input: s = "(]"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of parentheses only '()[]{}'.
//
// Related Topics String Stack 👍 13915 👎 647

package main

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := make(stackRunes, 0)
	for _, c := range stack {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			if stack.pop() != c {
				return false
			}
		}
	}
	return len(s) == 0
}

type stackRunes []rune

func (s *stackRunes) pop() rune {
	if len(*s) > 0 {
		last := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return last
	}
	return '\x00'
}

//leetcode submit region end(Prohibit modification and deletion)
