/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	var (
		buildTree func(val int)
		root *TreeNode
	)

	buildTree = func(val int) {
		

		a := &TreeNode{
			Val: val,
			Left: nil,
			Right: nil,
		}
	}

	return root
}
// @lc code=end

