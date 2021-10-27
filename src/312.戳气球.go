/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */
package leetcode

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n + 2)
	vals := make([]int, n + 2)

	vals[0], vals[n + 1] = 1, 1
	for i, v := range nums {
		vals[i+1] = v
	}

	for i := range dp {
		dp[i] = make([]int, n + 2)
	}

	for i := n - 1 ; i > -1; i-- {
		for j := i + 2; j < n + 2; j++ {
			for k := i + 1; k < j; k++ {
				sum := vals[i] * vals[j] * vals[k]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	
	return dp[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

