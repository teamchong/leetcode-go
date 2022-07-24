/*
 * @lc app=leetcode id=9 lang=rust
 *
 * [9] Palindrome Number
 */

// @lc code=start
impl Solution {
    /// We first check if the number is negative or if the number ends in 0. If either of these are true,
    /// we return false
    ///
    /// Arguments:
    ///
    /// * `x`: the number to be checked
    ///
    /// Returns:
    ///
    /// a boolean value.
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
