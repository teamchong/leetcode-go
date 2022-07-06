//Given an integer array nums, return an array answer such that answer[i] is
//equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
// You must write an algorithm that runs in O(n) time and without using the
//division operation.
//
//
// Example 1:
// Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
// Example 2:
// Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]
//
//
// Constraints:
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
//
//
// Follow up: Can you solve the problem in O(1) extra space complexity? (The
//output array does not count as extra space for space complexity analysis.)
// Related Topics Array Prefix Sum 👍 13063 👎 764

package main

//leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	right := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}

//leetcode submit region end(Prohibit modification and deletion)
