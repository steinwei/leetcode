/*
 * @lc app=leetcode.cn id=143 lang=typescript
 *
 * [143] 重排链表
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

/**
 Do not return anything, modify head in-place instead.
 */
function reorderList(head: ListNode | null): void {
    if(!head) return;
    if(!head.next) return;

    type listnode = ListNode | null

    let fast: listnode = head.next;
    let slow: listnode = head;

    while(fast && fast.next && slow) {
        fast = fast.next.next
        slow = slow.next
    }

    // const mid: listnode = slow
    let tmp: listnode = null
    if(slow && slow.next) {
        tmp = slow.next
        slow.next = null
    }

    let point: listnode = head
    const traversed = (tail: listnode) => {
        if(!tail) return tail
        traversed(tail.next)
        // console.log(point, tail)
        if(point) {
            tail.next = point.next;
            point.next = tail || null;
            point = point.next.next || null;
            // console.log(point, tail)
        }
        // head.next
    }

    traversed(tmp)
    // fuzhi

};
// @lc code=end

