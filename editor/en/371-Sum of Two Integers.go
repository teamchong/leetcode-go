//Given two integers a and b, return the sum of the two integers without using
//the operators + and -.
//
//
// Example 1:
// Input: a = 1, b = 2
//Output: 3
// Example 2:
// Input: a = 2, b = 3
//Output: 5
//
//
// Constraints:
//
//
// -1000 <= a, b <= 1000
//
// Related Topics Math Bit Manipulation 👍 2674 👎 3965

package main

//leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
