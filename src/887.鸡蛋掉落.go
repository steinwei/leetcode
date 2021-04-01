/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package leetcode

// @lc code=start
func superEggDrop(K int, N int) int {
	
	// memo
	memo := map[int]int{}

	var dp func(K,N int) int
	dp = func(K, N int) int {
		ans:=0
		// get memo
		if memo[N*100+K]  != 0 {
			return memo[K+N*100]
		}
		// base
		if K == 1 {
			ans = N
		} else if N == 0 {
			ans = 0
		} else {
			lo, hi := 1, N
			mid := (hi+lo)/2
			t1 := dp(K-1, mid - 1)
			t2 := dp(K, N-mid)
			// 二分法 更新基准值
			for lo + 1 < hi {
				if t1 > t2 {
					lo = mid
				} else if t1 < t2  {
					hi = mid
				} else {
					lo, hi = mid, mid
				}
			}
			ans = min(max(dp(K-1, lo - 1), dp(K, N-lo)), max(dp(K - 1, hi - 1), dp(K, N - hi))) + 1
		}
		// put memo
		memo[K+N*100] = ans

		return ans
	}

	return dp(K,N)
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
