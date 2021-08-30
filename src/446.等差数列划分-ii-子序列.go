/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */
package leetcode

import "sort"

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)

	// base

	// sort
	sort.Ints(nums)

	d := nums[1] - nums[0]

	ret, count := 0, 0

	for i := 2; i < n; i++ {
		
		if nums[i-1] - nums[i] == d {
			count++
		}
		ret += count
	}

	return ret
}
// @lc code=end

