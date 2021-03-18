/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */
package leetcode

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// @lc code=start
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	list := ListNode{Val: 0, Next: nil}
	cur := &list

	for head != nil {
		cur.Next = &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		cur = cur.Next
		head = head.Next
	}

	list = *list.Next

	return &list
}

// @lc code=end
