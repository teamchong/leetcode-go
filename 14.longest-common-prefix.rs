/*
 * @lc app=leetcode id=14 lang=rust
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
use std::cmp::min;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        for i in 0..strs[0].len() {
            if !strs[1..].iter().all(|s| s.get(i..=i) == strs[0].get(i..=i)) {
                return strs[0][..i].to_string();
            }
            
        }
        strs[0].to_string()
    }
}
// @lc code=end

