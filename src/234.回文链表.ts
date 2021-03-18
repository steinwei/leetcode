/*
 * @lc app=leetcode.cn id=234 lang=typescript
 *
 * [234] 回文链表
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

function isPalindrome(head: ListNode | null): boolean {
    let left : ListNode | null = head
    const travel = (right: ListNode | null): boolean => {
        if(!right) return true
        let res: boolean = travel(right.next)
        if(left) {
            res = res && left.val == right.val
            left = left.next
        }
        return res
    }
    return travel(head)
};
// @lc code=end

