/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
// If the current character is an opening bracket, push a corresponding closing bracket onto the stack
// as a signal for later. If the current character is a closing bracket, pop the top of the stack; if
// the popped character is the matching opening bracket for the current closing bracket, the brackets
// are balanced so far; if not, return false
func isValid(s string) bool {
	stack := make(Stack, 0)
	for _, c := range stack {
		switch c {
		case '(':
			stack.push(')')
		case '{':
			stack.push('}')
		case '[':
			stack.push(']')
		case ')', '}', ']':
			if stack.pop() != c {
				return false
			}
		}
	}
	return len(s) == 0
}

type Stack []rune

// A method that takes a pointer to a Stack and a rune value. It appends the rune value to the stack
// and returns the rune value.
func (s *Stack) push(val rune) rune { *s = append(*s, val) }

// Popping the last element of the stack.
func (s *Stack) pop() rune {
	if len(*s) > 0 {
		last := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return last
	}
	return '\x00'
}

// @lc code=end

