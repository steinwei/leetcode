/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package leetcode

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)

	ans := 0

	var dfs func(index ,ret int)
	dfs = func(index ,ret int) {
		if index > n - 1 {
			if ret == target {
				ans++	
			}
			return
		}

		x := nums[index] * -1
		y := nums[index]

		dfs(index+1, ret+x)
		dfs(index+1, ret+y)
	}

	dfs(0, 0)

	return ans
}
// @lc code=end

