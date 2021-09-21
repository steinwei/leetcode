/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	ans := [2]int{0, 0}

	var dfs func(node *TreeNode) [2]int
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		selected := node.Val + l[1] + r[1]
		notSelected := max(l[0], l[1]) + max(r[0], r[1])

		return [2]int{selected, notSelected}
	}

	ans = dfs(root)

	return max(ans[0], ans[1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

