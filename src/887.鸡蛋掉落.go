/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package leetcode

import (
	"math"
)

// @lc code=start
func superEggDrop(K int, N int) int {
	res := math.MaxInt64
	// binary-sort
	// var binary func(num int)
	var dp func(i, floor int) int
	dp = func(i, floor int) int {
		if K == 1 {
			return N
		}
		if N == 0 {
			return 0
		}
		if i > 1 && i < N {
			res = min(res,
				max(
					dp(i-1, floor-1),
					dp(i, N-floor),
				)+1)
		}
		return res
	}
	return dp(K, N)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
