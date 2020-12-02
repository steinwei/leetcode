/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package leetcode

// @lc code=start
func countPrimes(n int) int {
	ans := 0

	if n < 2 {
		return 0
	}

	count := 2
	for count < n {
		// if
		count++
	}

	return ans
}

// @lc code=end
