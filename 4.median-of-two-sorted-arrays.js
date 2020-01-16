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
var findMedianSortedArrays = function (nums1, nums2) {
    // merge into one array
    function merge(a, b) {
        return [].concat(a, b);
    }
    // sorted by quick-sort
    function quickSort(arr) {
        var len = arr.length;
        if (len <= 1) return arr;
        var pivotIndex = Math.floor(arr.length / 2);
        var pivot = arr.splice(pivotIndex, 1)[0];
        var left = [];
        var right = [];
        for (var i = 0; i < len; i++) {
            if (arr[i] < pivot) {
                left.push(arr[i]);
            } else {
                right.push(arr[i]);
            }
        }
        return quickSort(left).concat([pivot], quickSort(right));
    }
    function isOdd(num) {
        return num % 2 === 1;
    }
    // naturally
    function getMiddle(a, b) {
        var arr = merge(a, b);
        var arr1 = quickSort(arr);
        var len = arr1.length;
        var odd = isOdd(len);
        debugger;
        if (odd) {
            return (len + 1) / 2;
        } else {
            return (arr1[len / 2] + arr1[len / 2 + 1]) / 2;
        }
    }

    return getMiddle(nums1, nums2);
};
// @lc code=end

findMedianSortedArrays([1,2],[3,4]);