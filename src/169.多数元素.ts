/*
 * @lc app=leetcode.cn id=169 lang=typescript
 *
 * [169] 多数元素
 */

// @lc code=start
function majorityElement(nums: number[]): number {
    let ans = 0
    nums.reduce((pre: any, cur: number) : any=>{
        if(cur in pre) {
            pre[cur] ++
        } else {
            pre[cur] = 1
        }
        if(pre[cur] > Math.floor(nums.length / 2)) {
            ans = cur
        }
        return pre
    }, {})
    return ans
};
// @lc code=end

