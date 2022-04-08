/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
package leetcode

// @lc code=start
func rob(nums []int) int {
	var (
		n = len(nums)
		max func(x, y int) int
		dp func(start, end int) int
	)

	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp = func(start, end int) int {
		first, second := nums[start], max(nums[start], nums[start+1])
		for i := start+2; i <= end; i++ {
			first, second = second, max(second, first + nums[i])
		}
		return second
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(dp(0, n-2), dp(1, n-1))
}
// @lc code=end

