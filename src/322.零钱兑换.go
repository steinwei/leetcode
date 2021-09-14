/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package leetcode

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if v <= i {
				dp[i] = min(dp[i], dp[i - v] + 1)
			}
		}
	}

	if dp[amount] > amount {
		dp[amount] = -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

