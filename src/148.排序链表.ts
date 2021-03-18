/*
 * @lc app=leetcode.cn id=148 lang=typescript
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
// class ListNode {
//     val: number
//     next: ListNode | null
//     constructor(val?: number, next?: ListNode | null) {
//         this.val = (val===undefined ? 0 : val)
//         this.next = (next===undefined ? null : next)
//     }
// }

function sortList(head: ListNode | null): ListNode | null {
    type listnode = ListNode | null
    if(!head || !head.next ) return head
    let fast: listnode = head.next, slow: listnode = head
    // find mid
    while(fast.next && fast.next.next && slow) {
        fast = fast.next.next
        slow = slow.next
    }

    // reserve right
    let tmp: listnode = null
    // cut mid
    if(slow) {
        tmp = slow.next
        slow.next = null
    }

    let left:listnode = sortList(head)
    let right:listnode = sortList(tmp)
    
    let h:listnode = new ListNode(0)
    let res:listnode = h

    // console.log(left, right)
    // sort
    // loop
    while(left && right) {
        if( left.val < right.val) {
            h.next = left
            left = left.next
        }else {
            h.next = right
            right = right.next
        }
        h = h.next
    }
    // handle left nodes
    h.next = left ? left : right

    return res.next
};
// @lc code=end

