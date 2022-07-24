/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
package main

// We iterate through the array once, and for each element we check whether the consecutive sequence
// starting from that element is the longest sequence we've seen so far
func longestConsecutive(nums []int) int {
	longestStreak := 0
	if len(nums) > 0 {
		m := make(map[int]bool, len(nums))
		for _, num := range nums {
			m[num] = true
		}

		for num := range m {
			if !m[num-1] {
				streak := 1
				for i := num + 1; m[i]; i++ {
					streak++
				}
				if streak > longestStreak {
					longestStreak = streak
				}
			}
		}
	}
	return longestStreak
}

// @lc code=end
