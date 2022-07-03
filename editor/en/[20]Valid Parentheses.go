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
	stack := make([]rune, 0)
	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			if pop(&stack) != ch {
				return false
			}
		}
	}
	return len(stack) == 0
}
func pop(stack *[]rune) rune {
	if old := *stack; len(old) == 0 {
		return 0
	} else {
		lastIdx := len(old) - 1
		last := old[lastIdx]
		*stack = old[:lastIdx]
		return last
	}
}

//leetcode submit region end(Prohibit modification and deletion)
