/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int)  *ListNode {
	
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ;fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	// deleted
	slow.Next = slow.Next.Next

	return dummy.Next
}
// @lc code=end

