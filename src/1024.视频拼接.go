/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */
package leetcode

import "math"

// @lc code=start
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)

	const inf = math.MaxInt64 - 1

	for i := range dp {
		dp[i] = inf
	}

	// base case
	dp[0] = 0
	for i := 1; i < T+1; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}

	if dp[T] == inf {
		return -1
	}

	return dp[T]
}

// @lc code=end
