/*
 * @lc app=leetcode.cn id=1365 lang=typescript
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
function smallerNumbersThanCurrent(nums: number[]): number[] {
    const len = nums.length
    const ret = new Array<number>(len).fill(0)

    for (let i = 0; i < len; i++) {
        const cur = nums[i]
        for (let j = 0; j <  len; j++) {
            if(j == i) continue
            const last = nums[j]
            if(cur > last) {
                ret[i] ++
            }
        }        
    }

    return ret
};
// @lc code=end

