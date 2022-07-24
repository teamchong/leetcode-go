/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)
	for idx, num := range nums {
		if complementIdx, ok := complements[target-num]; ok {
			return []int{complementIdx, idx}
		}
		complements[num] = idx
	}
}

// @lc code=end

