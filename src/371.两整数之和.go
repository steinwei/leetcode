/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */
package leetcode

// @lc code=start
func getSum(a int, b int) int {
	var sum, curry int
	for b != 0 {
		sum = a ^ b
		curry = a & b << 1
		a = sum
		b = curry
	}

	return a
}
// @lc code=end

