/*
 * @lc app=leetcode.cn id=78 lang=typescript
 *
 * [78] 子集
 */

// @lc code=start
function subsets(nums: number[]): number[][] {
    // base case
    let res: number[][] = []

    // // 1. dsf
    const helper = (index: number, len: number, temp: number[], nums: number[]) => {
        // end
        if(index > nums.length){
            return
        }

        if(len > nums.length) {
            return
        }

        // res
        if(temp.length == len) {
            res.push(temp.concat())
        }

        // loop
        for (let i = index; i < nums.length; i++) {
            temp.push(nums[i])
            helper(i + 1, len + 1, temp, nums)
            temp.pop()
        }

    }

    helper(0, 0, [], nums)

    // 3. 借助栈实现dsf

    // const stack: number[][] = []
    // let start = 0
    // const end = nums.length

    // while(start < end){
    //     stack.push([nums[start++]])
    // }

    // while(stack.length!=0){
    //     const temp: number[] = stack.pop() || []

    //     for (let i = 0; i < nums.length; i++) {
    //         const clone = temp.concat()
    //         clone.push(nums[i])
            
    //     }


    // }

    // 2. 动规

    // const len = nums.length
    // // status
    // // 1. $len
    // // 2. $num
    // let $len = 1
    // // base case
    // res.push([])
    
    // while($len <= len){
    //     const stack: number[] = []
    //     const clone = nums.concat()
    //     let start = 0
    //     const end = len - $len
    //     let count = 0
    //     while(count < len){
            
    //         // loop
    //         for (let i = count; i < len; i++) {
    //             stack.push(clone[i])
    //             // res
    //             if($len == stack.length){
    //                 res.push(stack.concat())
    //                 stack.pop()
    //             }

    //             // end
    //             if(count == len - 1 && $len > 1) {
    //                 while(stack.length){
    //                     stack.pop()
    //                 }
    //                 count = start ++
    //                 if(start > end) {
    //                     start = 0
    //                     break
    //                 }
    //             }
    //         }
    //         count ++
    //     }
    //     $len++
    // }


    return res
};
// @lc code=end

subsets([1,2,3])