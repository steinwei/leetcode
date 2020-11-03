/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	count := len(prices)

	if count == 0 {
		return 0
	}

	dp := make([][2][2]int, count)

	// base case
	dp[0][0][0] = -prices[0]
	dp[0][0][1] = 0

	dp[0][1][0] = -prices[0]
	dp[0][1][1] = 0

	var findMax func(x, y int) int
	findMax = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < count; i++ {
		// 第一次成本
		dp[i][0][0] = findMax(dp[i-1][0][0], -prices[i])
		// 第一次利润
		dp[i][0][1] = findMax(dp[i-1][0][1], dp[i-1][0][0]+prices[i])
		// 第二次成本 = 第一次利润 - 当天价格
		dp[i][1][0] = findMax(dp[i-1][1][0], dp[i-1][0][1]-prices[i])
		// 第二次利润 = 第二次成本 + 当天价格
		dp[i][1][1] = findMax(dp[i-1][1][1], dp[i-1][1][0]+prices[i])
	}

	return dp[count-1][1][1]
}

// @lc code=end
