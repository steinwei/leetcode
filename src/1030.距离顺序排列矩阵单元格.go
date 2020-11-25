/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */
package leetcode

import "sort"

// @lc code=start
func allCellsDistOrder(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
