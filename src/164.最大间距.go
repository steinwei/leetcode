/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func maximumGap(nums []int) int {
	// base case
	if len(nums) <= 1 {
		return 0
	}

	ans := 0

	sort.Slice(nums, func(x, y int) bool {
		l, r := nums[x], nums[y]
		return r > l
	})

	// fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		l, r := nums[i-1], nums[i]
		ans = max(ans, r-l)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
