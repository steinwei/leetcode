/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */
package leetcode

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	cur := 9
	ans := 10
	for i := 0; i < n - 1; i++ {
		cur = (9-i)*cur
		ans += cur
	}

	return ans
}
// @lc code=end

