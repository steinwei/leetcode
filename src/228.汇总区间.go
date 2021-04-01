/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */
package leetcode

import "strconv"

// @lc code=start
func summaryRanges(nums []int) (ans []string) {

	for i, len := 0, len(nums); i < len; {
		left := i
		for i++; i < len && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}

	return ans
}

// @lc code=end
