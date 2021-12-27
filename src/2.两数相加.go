/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dumbhead = &ListNode{
			Val: -1,
			Next: nil,
		}
		l1cur = l1
		l2cur = l2
		carry = 0
		dumbtail = dumbhead
	)
	
	for l1cur != nil || l2cur != nil {
		n1, n2 := 0,0
		if l1cur != nil {
			n1 = l1cur.Val
			l1cur = l1cur.Next
		}

		if l2cur != nil {
			n2 = l2cur.Val
			l2cur = l2cur.Next
		}
		
		sum := (n1+n2+ carry) % 10
		carry = (n1+n2+ carry) / 10
		dumbtail.Next = &ListNode{
			Val: sum,
		}
		dumbtail = dumbtail.Next
	}

	if carry > 0 {
		dumbtail.Next = &ListNode{
			Val: carry,
		}
	}

	return dumbhead.Next
}
// @lc code=end

