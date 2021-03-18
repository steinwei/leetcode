/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package leetcode

// @lc code=start
func countPrimes(n int) int {
	ans := 0

	isPrime := make([]bool, n)
	for idx := range isPrime {
		isPrime[idx] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	return ans
}

// @lc code=end
