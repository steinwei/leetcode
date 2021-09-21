/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */
package leetcode

import "sort"

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])

	sorted := make([]int, n * m)
	index := 0

	for _, rows := range matrix {
		for _, cols := range rows {
			sorted[index] = cols
			index ++
		}
	}

	sort.Ints(sorted)

	return sorted[k - 1]
}
// @lc code=end

