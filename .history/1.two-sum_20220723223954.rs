/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut complements = HashMap::with_capacity(nums.len());
        for (i, num) in nums.iter().enumerate() {
            if let Some(j) = complements.get(&(target - num)) {
                return [*j, i];
            }
            complements.insert(num, i);
        }
        vec![]
    }
}
// @lc code=end

