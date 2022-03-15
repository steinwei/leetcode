/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */
package leetcode

import (
	"sort"
	"strconv"
)

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}

		for sy <= y {
			sy *= 10
		}

		// 这里的公式说实话，看不太懂。。
		return sy*x+y > sx*y+x
	})

	if nums[0] == 0 {
		return "0"
	}

	ans := []byte{}
	for _, v := range nums {
		ans = append(ans, strconv.Itoa(v)...)
	}

	return string(ans)
}
// @lc code=end

