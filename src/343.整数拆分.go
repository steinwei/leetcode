/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
package leetcode

// @lc code=start
func integerBreak(n int) int {
	var (
		f = make([]int, n+1)
		max func(x, y int) int
	)

	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 2; i <= n; i++ {
		maxNum := 0
		for j := 1; j < i; j++ {
			maxNum = max(maxNum, max(j*(i-j), f[i-j]*j))
		}
		f[i] = maxNum
	}

	return f[n]
}
// @lc code=end

