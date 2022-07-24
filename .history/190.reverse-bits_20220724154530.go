/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
package main

// We first swap the bits in the first and second half of the 32-bit number, then swap the bits in the
// first and second quarter, then the first and second eighth, and so on
func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xFF00FF00) >> 8) | ((num & 0x00FF00FF) << 8)
	num = ((num & 0xF0F0F0F0) >> 4) | ((num & 0x0F0F0F0F) << 4)
	num = ((num & 0xCCCCCCCC) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xAAAAAAAA) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// @lc code=end
