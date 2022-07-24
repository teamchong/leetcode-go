/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    /// We iterate through the array, and for each element, we check if the complement of the element
    /// exists in the hashmap. If it does, we return the indices of the element and its complement. If
    /// it doesn't, we add the element to the hashmap
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
        let mut complements: HashMap<i32, i32> = HashMap::with_capacity(nums.len());
        for (i, num) in nums.iter().enumerate() {
            if let Some(j) = complements.get(&(target - num)) {
                return vec![*j, i];
            }
            complements.insert(num, i);
        }
        vec![]
    }
}
// @lc code=end

