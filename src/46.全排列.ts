/*
 * @lc app=leetcode.cn id=46 lang=typescript
 *
 * [46] 全排列
 */

// @lc code=start
function permute(nums: number[]): number[][] {
    let res: number[][] = []
    const map:Map<number, boolean> = new Map()

    // $pos
    const backtracks = (path: number[] = []) => {
        // end
        if(path.length > nums.length) {
            return
        }
        // base case
        // res
        if(path.length == nums.length) {
            res.push(path.concat())
        }
        // loop
        for (let i = 0; i < nums.length; i++) {
            // skip
            if(map.get(nums[i])) continue
            // memo tracks
            map.set(nums[i], true)
            path.push(nums[i])
            backtracks(path)
            // pop tracks
            path.pop()
            map.set(nums[i], false)
        }
    }

    backtracks()

    return res
};
// @lc code=end

permute([1,2,3])