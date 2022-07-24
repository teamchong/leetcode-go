/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
package main

// We first check if the number is negative or ends with 0, then we reverse the number and compare it
// with the original number
func isPalindrome(x int) bool {
	if x%10 == 0 && x != 0 {
		return false
	}
	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x = x / 10
	}
	return x == reverted || x == reverted/10
}

// @lc code=end
