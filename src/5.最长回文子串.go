/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) string {
	ans := ""

	n:= len(s)

	if n==1 {
		return s
	}

	// 建阵，将问题转化为一个结果矩阵
	dp:= make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// 用来比较最长的长度，记录长度的指针
	maxLen := 1
	// 首位，用来记录首位的指针
	begin := 0

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			
			// L = j - i + 1 ,j 代表尾部
			j := L + i - 1

			// 超出右边界则跳出
			if j >= n || j < 0 {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 三个以内首尾相等基本相当于回文了
				if j - i < 3 {
					dp[i][j] = true
				} else {
					// move
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 
			if dp[i][j] && j - i + 1 > maxLen{
				// update maxLen
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	// 通过指针取对应的子串
	ans = s[begin: begin+maxLen]
	
	return ans
}
// @lc code=end

 