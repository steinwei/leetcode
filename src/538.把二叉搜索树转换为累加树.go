/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(p *TreeNode) 
	dfs = func(p *TreeNode) {
		if p != nil {
			dfs(p.Right)
			sum += p.Val
			p.Val = sum
			dfs(p.Left)
		}
	}

	dfs(root)

	return root
}
// @lc code=end

