/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
package main

//leetcode submit region begin(Prohibit modification and deletion)
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
