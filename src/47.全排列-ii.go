/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package leetcode

import "sort"

// @lc code=start
func permuteUnique(nums []int) (ans [][]int) {
	var (
		dfs func(start int, tmp []int, seen []bool)
		n = len(nums)
	)

	// 题目有写任意顺序，所以这里需要先排序
	sort.Ints(nums)

	dfs = func(start int, tmp []int, seen []bool) {
		if start == n {
			// 这里需要 clone 不然，会导致数据出错，复用到上个遍历的 tmp 数组
			ans = append(ans, append([]int{}, tmp...))
			return
		}

		for i,v := range nums {
			// 其实是条件限制少了，思路是没问题的，要排除掉元素相同的情况
			if seen[i] || i>0 && !seen[i-1] && nums[i-1]==v {
				continue
			}
			tmp = append(tmp, nums[i])
			seen[i] = true
			dfs(start+1,tmp, seen)
			seen[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0,[]int{}, make([]bool, n))

	return
}
// @lc code=end

