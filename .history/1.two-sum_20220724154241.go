/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
// We iterate through the array, adding each element's value and its index to the table. Then, before
// we've added the second element, we check if the table already contains the complement we're looking
// for. If it does, we've found a solution and we return immediately
package main

func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)
	for idx, num := range nums {
		if complementIdx, ok := complements[target-num]; ok {
			return []int{complementIdx, idx}
		}
		complements[num] = idx
	}
	return nil
}

// @lc code=end
