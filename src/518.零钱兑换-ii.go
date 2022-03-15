/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */
package leetcode

// @lc code=start
func change(amount int, coins []int) int {
	var (
		dp = make([]int, amount+1)
	)

	dp[0] = 1
	for _, v := range coins {
		for i := v; amount - i >= 0; i++ {
			dp[i] += dp[i-v]
		}
	}

	return dp[amount]
}
// @lc code=end

