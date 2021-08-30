/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	
	stack := []int{}

	var inorder func(node *TreeNode) 
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		stack = append(stack, node.Val)

		inorder(node.Right)
	}

	inorder(root)
	
	return stack[k-1]
}
// @lc code=end

