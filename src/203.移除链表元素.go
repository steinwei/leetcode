/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */
package leetcode


type ListNode struct {
	Val int
	Next *ListNode
}
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	list := &ListNode{
		Val: 0,
		Next: head,
	}
	prev, cur := list, head

	for cur!=nil {
		if cur.Val != val {
			prev.Next = cur			
			prev = prev.Next
		}
		if cur.Next == nil {
			prev.Next = nil
		}
		cur = cur.Next
	}

	return list.Next
}
// @lc code=end

