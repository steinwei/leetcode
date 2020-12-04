/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package leetcode

// @lc code=start
func countPrimes(n int) int {
	ans := 0

	var isPrime = func(num int) bool {
		for i := 0; i*i < num; i++ {
			if num&i == 0 {
				return false
			}
		}
		return true
	}

	count := 2
	for count < n {
		if isPrime(count) {
			ans++
		}
	}

	return ans
}

// @lc code=end
