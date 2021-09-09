/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package leetcode

// @lc code=start
func canJump(nums []int) bool {

	n := len(nums)

	dp := make([]int, n)

	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		if i <= dp[i-1] {
			curr := nums[i] + i
			dp[i] = max(curr, dp[i-1])
		}
	}

	return dp[n-1] >= n - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
// @lc code=end

