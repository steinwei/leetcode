/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */
package leetcode

// @lc code=start
func countDigitOne(n int) int {
	var (
		f = make([]int, n+1)	
	)

	f[0] = 0
	for i := 1; i <= n; i++ {
		if i >= 10 {
			cur := i
			for cur>0 {
				if f[cur]>0 {
					f[i]+= f[cur]
					break
				}
				if cur%10 == 1 {
					f[i] += 1
				}
				cur/=10
			}
		} else {
			if i%10 == 1 {
				f[i] += 1
			}
		}
		f[i] += f[i-1]
	}

	return f[n]
}
// @lc code=end

