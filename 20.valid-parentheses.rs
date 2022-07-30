/*
 * @lc app=leetcode id=20 lang=rust
 *
 * [20] Valid Parentheses
 */

// @lc code=start
impl Solution {
    pub fn is_valid(s: String) -> bool {
        if s.is_empty() {
            return true;
        }
        let mut stack = Vec::new();
        for c in s.chars() {
            match c {
                '{' => stack.push('}'),
                '(' => stack.push(')'),
                '[' => stack.push(']'),
                '}'|')'|']' if Some(c) != stack.pop() => return false,
                _ => (),
            }
        }   
        return stack.is_empty();
    }
}
// @lc code=end

