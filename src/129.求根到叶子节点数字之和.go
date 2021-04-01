/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 */
package leetcode

// @lc code=start
/**
* Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sumNumbers(root *TreeNode) int {
	//
	return dfs(root, 0)

}

func dfs(head *TreeNode, prevSum int) int {
	if head == nil {
		return 0
	}

	sum := head.Val + prevSum*10
	// fmt.Println(sum, prevSum)

	if head.Left == nil && head.Right == nil {
		return sum
	}

	return dfs(head.Left, sum) + dfs(head.Right, sum)
}

// @lc code=end
