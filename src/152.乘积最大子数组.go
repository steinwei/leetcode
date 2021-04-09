/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package leetcode

// @lc code=start
func maxProduct(nums []int) int {
	n:= len(nums)
	dp := make([]int, n)

	dp[0] = 0
	for i := 1; i < n; i++ {
		sum := max(nums[i] * nums[i-1], nums[i])
		dp[i] = max(dp[i-1], sum)
	}

	return dp[n-1]
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}
// @lc code=end

