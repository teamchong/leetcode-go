/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
package main

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

// @lc code=end
