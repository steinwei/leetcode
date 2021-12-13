/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		dumbhead = &ListNode{
			Val: -1,
			Next: head,
		}
		lprev = dumbhead
		reverse func(start *ListNode)
	)

	reverse = func(start *ListNode){
		curr := start
		var prev *ListNode
		for curr!=nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
	}

	for i := 0; i < left-1; i++ {
		lprev = lprev.Next
	}
	
	l := lprev.Next
	r := lprev
	for i := 0; i < right-left+1; i++ {
		r = r.Next
	}
	rnext := r.Next

	// 截断链表
	lprev.Next = nil
	r.Next = nil
	reverse(l)

	// 连接链表
	lprev.Next=r
	l.Next=rnext

	return dumbhead.Next
}
// @lc code=end

