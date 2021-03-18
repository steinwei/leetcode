/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)

	// base case
	dp[0][0] = -prices[0]

	var findMax func(x, y int) int
	findMax = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < n; i++ {
		dp[i][0] = findMax(dp[i-1][0], -prices[i])
		dp[i][1] = findMax(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
}

// @lc code=end
