/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */
package leetcode

import "math/rand"

// @lc code=start
type RandomizedSet struct {
	nums []int
	indicates map[int]int
}


func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indicates[val]; ok {
		return false
	}
	this.indicates[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	id, ok :=  this.indicates[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	this.indicates[this.nums[id]] = id
	this.nums = this.nums[:last]
	delete(this.indicates, val)
	return true
}


func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

