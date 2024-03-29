/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

package leetcode

import "math"

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j * j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}

	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

