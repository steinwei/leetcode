/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	
	ans := [][]int{}
	
	if root == nil {
		return ans
	}
	
	tmp := []*TreeNode{root}

	for i := 0; len(tmp) > 0; i++ {
		ans = append(ans, []int{})
		temp := []*TreeNode{}
		for j := 0; j < len(tmp); j++ {
			node := tmp[j]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		tmp = temp
	}

	return ans
}
// @lc code=end

