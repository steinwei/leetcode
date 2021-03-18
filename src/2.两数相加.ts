/*
 * @lc app=leetcode.cn id=2 lang=typescript
 *
 * [2] 两数相加
 */

// @lc code=start
// Definition for singly-linked list.
// class ListNode {
//     val: number
//     next: ListNode | null
//     constructor(val?: number, next?: ListNode | null) {
//         this.val = (val===undefined ? 0 : val)
//         this.next = (next===undefined ? null : next)
//     }
// }

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {

    type listnode = ListNode | null

    let head:listnode = null, tail:listnode = null

    let carry = 0

    while(l1 || l2){

        const n1 = l1 ? l1.val : 0
        const n2 = l2 ? l2.val : 0

        const sum = n1 + n2 + carry

        if(!head){
            head = tail = new ListNode(sum % 10)
        } else if(tail) {
            tail.next = new ListNode(sum % 10)
            tail = tail.next
        }
        carry = Math.floor( sum / 10)

        if(l1) {
            l1 = l1.next
        }

        if(l2) {
            l2 = l2.next
        }

    }
    if(carry && tail) {
        tail.next = new ListNode(carry)
    }

    return head
};
// @lc code=end

