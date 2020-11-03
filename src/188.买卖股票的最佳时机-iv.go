/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package leetcode

// @lc code=start
func maxProfit(k int, prices []int) int {
	count := len(prices)

	if count == 0 {
		return 0
	}

	if k > count/2 {
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

	dp := make([][][]int, count)

	for i := 0; i < count; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	var findMax func(x, y int) int
	findMax = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < count; i++ {
		for j := 1; j <= k; j++ {
			// base case
			if i == 0 {
				// 第一次成本
				dp[i][j][0] = -prices[0]
				// 第一次利润
				dp[i][j][1] = 0
				continue
			}
			// 第二次成本 = 第一次利润 - 当天价格
			dp[i][j][0] = findMax(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i])
			// 第二次利润 = 第二次成本 + 当天价格
			dp[i][j][1] = findMax(dp[i-1][j][1], dp[i-1][j][0]+prices[i])

		}
	}

	return dp[count-1][k][1]
}

// @lc code=end
