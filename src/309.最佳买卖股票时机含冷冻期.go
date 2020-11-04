/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)

	prev := 0

	for i := 0; i < n; i++ {
		cost := prices[i]

		if i == 0 {
			dp[i][0] = -cost
			dp[i][1] = 0
			continue
		} else if i >= 2 {
			prev = dp[i-2][1]
		}

		dp[i][0] = max(dp[i-1][0], prev-cost)
		dp[i][1] = max(dp[i-1][1], cost+dp[i][0])

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
