/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{0, head}

	slow, fast := head, head
	n := 1

	for fast.Next != nil {
		n ++
		fast = fast.Next
	}

	// 闭合为环
	fast.Next = head

	for i := 0; i < n-k % n - 1; i++ {
		slow = slow.Next
	}
	
	dummy.Next = slow.Next
	slow.Next = nil

	return dummy.Next
}
// @lc code=end

