/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
package main

// For each number i from 1 to n, dp[i] = dp[i / 2] + i % 2
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}

// @lc code=end
