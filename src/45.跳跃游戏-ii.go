/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package leetcode

// @lc code=start
func jump(nums []int) int {
	var (
		n = len(nums)
		maxn = 0
		end = 0
		steps = -1
	)

	for i := 0; i<n; i++ {
		curr := i+nums[i]
		if curr >= n-1 {
			curr = n-1
		}
		maxn = max(curr, maxn)
		if i == end {
			end = maxn
			steps ++
		}
	}

	return steps
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

