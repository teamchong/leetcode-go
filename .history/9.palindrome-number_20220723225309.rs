/*
 * @lc app=leetcode id=9 lang=rust
 *
 * [9] Palindrome Number
 */

// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x % 10 == 0 && x != 0 {
            return false;
        }
        let mut right = 0;
        while x > right {
            right = right * 10 + x % 10;
            x /= 10;
        }
        x == right || x == right / 10
    }
}
// @lc code=end

