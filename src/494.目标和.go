/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package leetcode

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// n := len(nums)

	// ans := 0

	// var dfs func(index ,ret int)
	// dfs = func(index ,ret int) {
	// 	if index > n - 1 {
	// 		if ret == target {
	// 			ans++	
	// 		}
	// 		return
	// 	}

	// 	x := nums[index] * -1
	// 	y := nums[index]

	// 	dfs(index+1, ret+x)
	// 	dfs(index+1, ret+y)
	// }

	// dfs(0, 0)

	// return ans

	// 下面用 dp 来作答

	var (
		sum = 0
	)
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff & 1 == 1 {
		return 0
	}
	neg := diff >> 1
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, v := range nums {
		for i := neg; i - v >= 0; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[neg]
}
// @lc code=end

