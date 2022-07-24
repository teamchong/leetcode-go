/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
package main

// If the number is already in the map, return true. Otherwise, add it to the map.
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

// @lc code=end
