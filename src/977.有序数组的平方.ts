/*
 * @lc app=leetcode.cn id=977 lang=typescript
 *
 * [977] 有序数组的平方
 */

// @lc code=start
function sortedSquares(A: number[]): number[] {
    return A.map(num=>Math.pow(num, 2)).sort((a, b) => a - b)
};
// @lc code=end

