/*
 * @lc app=leetcode id=13 lang=rust
 *
 * [13] Roman to Integer
 */

// @lc code=start
impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        s.chars().fold(0, |acc, c| {
            acc + match c {
                'I' => 1,
                'V' if acc % 10 == 1 => 3,
                'V' => 5,
                'X' if acc % 10 == 1 => 8,
                'X' => 10,
                'L' if acc % 100 == 10 => 30,
                'L' => 50,
                'C' if acc % 100 == 10 => 80,
                'C' => 100,
                'D' if acc % 1000 == 100 => 300,
                'D' => 500,
                'M' if acc % 1000 == 100 => 800,
                'M' => 1000,
                _ => 0
            }
        })
    }
}
// @lc code=end

