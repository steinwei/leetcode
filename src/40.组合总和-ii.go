/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
package leetcode

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)

	var dfs func(pos, sum int, temp []int)
	dfs = func(pos, sum int, temp []int) {
		if sum == target {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}

		if sum > target {
			return
		}

		for i := pos; i < len(candidates); i++ {
			if pos <= i - 1 && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, sum + candidates[i], temp)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, 0, []int{})

	return
}
// @lc code=end

