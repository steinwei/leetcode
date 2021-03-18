/*
 * @lc app=leetcode.cn id=11 lang=typescript
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
function maxArea(height: number[]): number {
    let res =  0

    let l = 0, r = height.length - 1

    while(l<=r){
        // end?
        const min = Math.min(height[l], height[r])
        const ans = min * (r - l)

        // res?
        res = Math.max(res, ans)

        // loop?
        if(min == height[l]){
            l++
        } else {
            r--
        }

    }

    return res
};
// @lc code=end

