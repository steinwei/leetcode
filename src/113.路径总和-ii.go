/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}

	var dfs func(head *TreeNode)
	dfs = func(head *TreeNode) {
		ans = append(ans, head.Val)
	}

	return
}
// @lc code=end

