/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    /// We iterate through the vector, and for each element, we check if the complement of the element
    /// exists in the hashmap. If it does, we return the indices of the element and its complement. If
    /// it doesn't, we insert the element into the hashmap
    /// 
    /// Arguments:
    /// 
    /// * `nums`: Vec<i32> - a vector of integers
    /// * `target`: the target sum
    /// 
    /// Returns:
    /// 
    /// A vector of two integers.
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut complements = HashMap::with_capacity(nums.len());
        for (i, num) in nums.iter().enumerate() {
            let complement = target - num;
            if complements.contains_key(&complement) {
                return vec![complements[&complement], i as i32];
            }
            complements.insert(num, i as i32);
        }
        vec![]
    }
}
// @lc code=end

