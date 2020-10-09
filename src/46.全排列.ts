/*
 * @lc app=leetcode.cn id=46 lang=typescript
 *
 * [46] 全排列
 */

// @lc code=start
function permute(nums: number[]): number[][] {
    let res: number[][] = []

    const helper = (index: number, temp: number[], nums: number[]) => {
        // end
        if(index > nums.length) {
            return
        }
        // base case
        // res
        if(temp.length == nums.length) {
            res.push(temp.concat())
        }
        // loop
        for (let i = index; i < nums.length; i++) {
            temp.push(nums[i])
            helper(i + 1, temp, nums)
            temp.pop()
        }
    }

    helper(0, [], nums)

    return res
};
// @lc code=end

