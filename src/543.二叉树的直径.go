/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		L := dfs(node.Left)
		R := dfs(node.Right)

		ans = max(L+R, ans)

		return max(L, R) + 1
	}

	dfs(root)

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

