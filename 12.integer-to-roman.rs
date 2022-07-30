/*
 * @lc app=leetcode id=12 lang=rust
 *
 * [12] Integer to Roman
 */

// @lc code=start
impl Solution {
    pub fn int_to_roman(num: i32) -> String {
        let (mut result, mut num) = (String::new(), num);
        let mut values = vec![1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        let mut symbols = vec!["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        for i in 0..values.len() {
            let count = num / values[i];
            if count > 0 {
                result.push_str(&symbols[i].repeat(count as usize));
                num -= count * values[i];
            }
        }
        result
    }
}
// @lc code=end

