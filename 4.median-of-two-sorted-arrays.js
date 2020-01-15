/*
 * @lc app=leetcode id=4 lang=javascript
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    function merge(nums1,nums2){
        const arr = [];
        arr.concat(...nums1,...nums2);
        return arr;
    }
    function sort(nums){
        var len =  nums.length;
        for(var i = len;i>0;i--){
            nums[i]
        }
    }
    function getMiddle(){

    }

    return getMiddel();
};
// @lc code=end

