/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */
package leetcode

import "math/rand"

// @lc code=start
type RandomizedCollection struct {
	idx  map[int]map[int]struct{}
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: map[int]map[int]struct{}{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (r *RandomizedCollection) Insert(val int) bool {
	ids, has := r.idx[val]
	if !has {
		ids = map[int]struct{}{}
		r.idx[val] = ids
	}
	ids[len(r.nums)] = struct{}{}
	r.nums = append(r.nums, val)
	return !has
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (r *RandomizedCollection) Remove(val int) bool {
	ids, has := r.idx[val]
	if !has {
		return false
	}
	var i int
	// 找到最近的一个 map
	for id := range ids {
		i = id
		break
	}
	n := len(r.nums)
	// 更新数组的值 模拟栈 出栈顶 改变栈内数值
	r.nums[i] = r.nums[n-1]
	// 删除索引
	delete(ids, i)
	// 删除map中索引
	delete(r.idx[r.nums[i]], n-1)
	// 更新map
	if i < n-1 {
		r.idx[r.nums[i]][i] = struct{}{}
	}
	// 删除内嵌map索引
	if len(ids) == 0 {
		delete(r.idx, val)
	}
	// 更新数组长度
	r.nums = r.nums[:n-1]
	return true
}

/** Get a random element from the collection. */
func (r *RandomizedCollection) GetRandom() int {
	// 从nums数组中随机返回 这个简单
	return r.nums[rand.Intn(len(r.nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
