/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
package main

// We use XOR to add two numbers without carry, and AND to carry
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

// @lc code=end
