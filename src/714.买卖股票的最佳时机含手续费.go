/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */
package leetcode

// @lc code=start
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)

	if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		// base case
		if i == 0 {
			dp[i][0] = -prices[i] - fee
			dp[i][1] = 0
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		// fmt.Println(dp[i][0], dp[i][1])
	}

	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
