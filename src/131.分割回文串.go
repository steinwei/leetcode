/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package leetcode

// @lc code=start
func partition(s string) (ans [][]string) {
	n := len(s)

	f := make([][]bool, n)

	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	// 动规预处理 「回文串」
	for i := n - 1; i > -1; i-- {
		for j := i+1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	res:=[]string{}

	// 回溯模版套入，不解释
	var dfs func(int)
	dfs = func(i int) {
		if(i == n) {
			ans = append(ans, append([]string(nil), res...))
			return
		}

		for j := i; j < n; j++ {
			// 先判断是否回文串，再作处理
			if(f[i][j]) {
				res = append(res, s[i: j+1])
				dfs(j+1)
				res = res[:len(res)-1]
			}
		}
	}

	dfs(0)

	return
}
// @lc code=end

