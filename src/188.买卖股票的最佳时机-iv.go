/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package leetcode

import "math"

// @lc code=start
func maxProfit(k int, prices []int) int {
	count := len(prices)

	if count == 0 {
		return 0
	}

	if k > count/2 {
		return maxProfit2(prices)
	}

	dp_i_0, dp_i_1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i < k+1; i++ {
		dp_i_0[i] = 0
		dp_i_1[i] = math.MinInt64
	}
	for _, v := range prices {
		for i := k; i >= 1; i-- {
			dp_i_0[i] = max(dp_i_0[i], dp_i_1[i]+v)
			dp_i_1[i] = max(dp_i_1[i], dp_i_0[i-1]-v)
		}
	}
	return dp_i_0[k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, -1<<31
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, temp-prices[i])
	}
	return dp_i_0
}

// @lc code=end
