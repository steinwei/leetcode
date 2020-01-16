/*
 * @lc app=leetcode id=1 lang=javascript
 *
 * [1] Two Sum
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var i,j,len = nums.length;
    var hash = {};
    for (i=0;i<len;i++) {
        j = target-nums[i];
        if (hash[j] !== undefined)
            return [i,hash[j]];
        else 
            hash[nums[i]] = i;
    }
};


