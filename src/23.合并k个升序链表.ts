/*
 * @lc app=leetcode.cn id=23 lang=typescript
 *
 * [23] 合并K个升序链表
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
 * solution
 * 1. find min head, skip to next
 * 2. compare each listnode value and find min to appent to listnode
 * 3. if one ended and others doesnt, travel others and repeat steps 2
 * 4. repeat steps 2 and 3 util all run through
 * 5. get sorted listnode
 * @param lists 
 */
function mergeKLists(lists: Array<ListNode | null>): ListNode | null {
    type listnode = ListNode | null
    let head:listnode = null, tail:listnode = null
    let min = Infinity
    let point:number = 0
    const k = lists.length

    while (true) {
        for (let i = 0; i < k; i++) {
            const l = lists[i]
            if(!l) continue
            min = Math.min(l.val, min)
            if(min == l.val) {
                point = i
            }
        }
        // end
        if(!lists[point]) break

        if(!head) {
            head = tail = new ListNode(min)
        } else if(tail){
            tail.next = new ListNode(min)
            tail = tail.next
        }

        // reset
        min = Infinity

        // console.log(lists[point])
        // travel
        lists[point] = (lists[point]|| new ListNode(Infinity)).next

    }

    return head
};
// @lc code=end

