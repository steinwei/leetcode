/*
 * @lc app=leetcode.cn id=111 lang=typescript
 *
 * [111] 二叉树的最小深度
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

// class TreeNode {
//     val: number
//     left: TreeNode | null
//     right: TreeNode | null
//     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
//         this.val = (val===undefined ? 0 : val)
//         this.left = (left===undefined ? null : left)
//         this.right = (right===undefined ? null : right)
//     }
// }


function minDepth(root: TreeNode | null): number {

    if(!root) {
        return 0
    }
    if(!root.left && !root.right) {
        return 1
    }
    if(!root.left) {
        return minDepth(root.right) + 1
    }
    if(!root.right) {
        return minDepth(root.left) + 1
    }

    let count = Math.min(minDepth(root.left), minDepth(root.right))

    console.log(count)

    return count + 1
};
// @lc code=end

