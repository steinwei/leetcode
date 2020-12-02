/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */
package leetcode

import "sort"

// @lc code=start
func largestPerimeter(A []int) int {
	sort.Ints(A)
	// c - a < b
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}

	return 0
}

// @lc code=end
