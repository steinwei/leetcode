/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */
package leetcode

import "sort"

// @lc code=start
func findDuplicate(nums []int) int {

	sort.Ints(nums)
	
	n := len(nums)

	ans := -1

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			ans = nums[i]
			break
		}
	}

	return ans
}
// @lc code=end

