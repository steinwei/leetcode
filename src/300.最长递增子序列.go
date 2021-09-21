/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package leetcode

// @lc code=start
func lengthOfLIS(nums []int) int {

	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	dp[0] = 1
	maxn := dp[0]

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j]  < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		maxn = max(maxn, dp[i])
	}

	return maxn
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

