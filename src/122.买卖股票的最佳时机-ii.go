/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	count := len(prices)

	if count == 0 {
		return 0
	}

	dp := make([][2]int, count)

	// base case
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	var findMax func(x, y int) int
	findMax = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < count; i++ {
		dp[i][0] = findMax(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = findMax(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[count-1][1]
}

// @lc code=end
