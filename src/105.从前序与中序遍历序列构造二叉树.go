/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)

	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
		Left: nil,
		Right: nil,
	}

	i := 0
	for ; i < n; i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1: len(inorder[:i]) + 1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i]) + 1:], inorder[i+1:])

	return root
}
// @lc code=end

