/*
 * @lc app=leetcode.cn id=4 lang=typescript
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    // let len1 = nums1.length
    // let len2 = nums2.length
    // let curr = 0, prev = 0
    // let res = 0
    // let i = 0, j = 0, count = 1
    // const mid = (len1+len2)/2
    // // loop condition?
    // while(count<len1+len2){
    //     // end condition?

    //     // result condition?
    //     if(count==mid + 1){
    //         break
    //     }else if(count==mid+0.5){
    //         res = curr
    //         break
    //     }

    //     // loop condition?
    //     curr = Math.max(nums1[i], nums2[j])
    //     if(curr == nums1[i]){
    //         prev = nums2[j]
    //         j++
    //     } else {
    //         prev = nums1[i]
    //         i++
    //     }

    //     count++
    // }
    // return res

    const arr:number[] = nums1.concat(nums2)
    arr.sort((l,r)=>l-r)
    // console.log(arr)
    const len = nums1.length + nums2.length
    if(len&1){
        const mid = (len-1)/2 
        return arr[mid]
    }else{
        const mid = len / 2 - 1
        return (arr[mid]+arr[mid+1])/2
    }
};
// @lc code=end
