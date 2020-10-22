/*
 * @lc app=leetcode.cn id=238 lang=typescript
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
function productExceptSelf(nums: number[]): number[] {
    const len = nums.length;
    const ans = new Array<number>(len);

    ans[0] = 1;
    for (let i = 1; i < len; i++) { 
        ans[i] = ans[i - 1] * nums[i - 1];
    }

    let R = 1;
    for (let i = len - 1; i > -1; i--) { 
        ans[i] = ans[i] * R;
        R *= nums[i];
    }

    return ans
};
// @lc code=end

