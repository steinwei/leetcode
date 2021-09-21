/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	list := []*TreeNode{root}
	var prev *TreeNode
	for len(list) > 0 {
		curr := list[len(list)-1]
		list = list[:len(list)-1]

		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}

		left := curr.Left
		right := curr.Right

		if right != nil {
			list = append(list, right)
		}

		if left != nil {
			list = append(list, left)
		}

		prev = curr
	}

	return
}
// @lc code=end

