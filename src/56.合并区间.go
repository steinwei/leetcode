/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package leetcode

import "sort"

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}

	res = append(res, prev)

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

