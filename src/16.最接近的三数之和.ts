/*
 * @lc app=leetcode.cn id=16 lang=typescript
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
function threeSumClosest(nums: number[], target: number): number {

    let ans = 0
    let min = Infinity
    // sort
    nums.sort((a, b)=> a - b)
    console.log(nums)

    const len = nums.length - 1

    // double point start tail and cur
    for (let i = 0; i <= len;) {
        let head = i + 1
        let tail = len
        while(head < tail) {
            const first = nums[head]
            const second = nums[tail]
            const sum = first + second + nums[i]
            const result = sum - target
            // min abs
            // console.log('%s:%d', 'result', Math.abs(result), min)
            if( Math.abs(result) <= min) {
                ans = sum
            }
            min = Math.min(min, Math.abs(result))
            if(result<=0){
                while(nums[head]==nums[++head]){}
            }else if(result>0){
                while(nums[tail]==nums[--tail]){}
            }else{
                break
            }
        }
        while(nums[i]==nums[++i]){}
        // get sum
    }

    return ans
};
// @lc code=end

