/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */
package leetcode

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32

	var inorder func(node *TreeNode) int
	inorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := max(inorder(node.Left), 0)
		r := max(inorder(node.Right), 0)

		sum := node.Val + l + r

		ans =  max(sum, ans)

		return node.Val + max(l, r)
	}

	inorder(root)

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
// @lc code=end

