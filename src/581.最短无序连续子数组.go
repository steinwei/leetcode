/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */
package leetcode

import "sort"

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	ans := 0
	n := len(nums)

	clone := make([]int, n)
	copy(clone, nums)
	sort.Ints(clone)

	i, j := 0, n-1
	for ; i <= j && nums[i] == clone[i]; i++ {}
	for ; j >= i && nums[j] == clone[j]; j-- {}

	ans = j - i + 1

	return ans
}
// @lc code=end

