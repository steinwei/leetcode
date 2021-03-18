/*
 * @lc app=leetcode.cn id=144 lang=typescript
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */


//  class TreeNode {
//      val: number
//      left: TreeNode | null
//      right: TreeNode | null
//      constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
//          this.val = (val===undefined ? 0 : val)
//          this.left = (left===undefined ? null : left)
//          this.right = (right===undefined ? null : right)
//      }
//  }
 

function preorderTraversal(root: TreeNode | null): number[] {
    const ret: number[] = []
    if(!root) return ret
    let head: TreeNode | null = root
    let stack: TreeNode[]= []
    while(stack.length != 0 || head != null) {
        while(head != null) {
            ret.push(head.val)
            stack.push(head)
            head = head.left
        }
        if(stack.length != 0) {
            head = stack.pop() || null
            head = head ? head.right : null
        }
    }
    return ret
};
// @lc code=end

