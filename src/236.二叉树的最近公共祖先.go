/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	L := lowestCommonAncestor(root.Left, p, q)
	R := lowestCommonAncestor(root.Right, p, q)

	if L != nil && R != nil {
		return root
	}

	if L != nil {
		return L
	}

	return R
}
// @lc code=end

