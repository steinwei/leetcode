/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}

		seen[head] = struct{}{}
		head = head.Next
	}

	return nil

}
// @lc code=end

