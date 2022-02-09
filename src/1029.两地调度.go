/*
 * @lc app=leetcode.cn id=1029 lang=golang
 *
 * [1029] 两地调度
 */
package leetcode

import "sort"

// @lc code=start
func twoCitySchedCost(costs [][]int) int {
	var (
		ans = 0
		n = len(costs)
		mid = n / 2
	)

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0] - costs[i][1] < costs[j][0] - costs[j][1]
	})

	for i := 0; i < n/2; i++ {
		ans += costs[i][0] + costs[i+mid][1]
	}

	return ans
}
// @lc code=end

