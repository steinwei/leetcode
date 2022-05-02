/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */
package leetcode

import "math"

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	var (
		f = make([]int, n+1)
		m = len(primes)
		pointers = make([]int, m)
		nums = make([]int, m)
		min func(x, y int) int
	)

	min = func(x, y int) int {
		if x<y {
			return x
		}
		return y
	}

	for i := range nums {
		nums[i] = 1
	}
	
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range pointers {
			minNum = min(minNum, nums[j])
		}
        f[i] = minNum
		for j := range nums {
			if nums[j] == minNum {
				pointers[j] ++
				nums[j] = f[pointers[j]] * primes[j]
			}
		}
	}

	return f[n]
}
// @lc code=end

