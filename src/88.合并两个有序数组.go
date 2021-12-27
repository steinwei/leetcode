/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package leetcode

import "sort"

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = nums1[0: m]
	nums2 = nums2[0: n]

	nums1 = append(nums1, nums2...)

	// todo: 这里使用排序算法即可
	sort.Ints(nums1)

	return
}
// @lc code=end

