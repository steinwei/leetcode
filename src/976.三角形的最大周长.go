/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */
package leetcode

import "sort"

// @lc code=start
func largestPerimeter(A []int) int {
	ans := 0

	// a + b > c
	// c - a < b

	sort.Slice(A, func(x, y int) bool {
		return A[x] > A[y]
	})

	lo, hi := 0, len(A)-1
	for hi > lo {
		left, right := A[lo], A[hi]
		if right-left >= left {
			hi--
		} else {
			lo++
		}
	}

	return ans
}

// @lc code=end
