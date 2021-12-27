/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package leetcode

// @lc code=start
func generateParenthesis(n int) (ans []string) {
	var (
		dfs func(lcnt, rcnt int, s string)
	)

	dfs = func(lcnt, rcnt int, s string) {
		if lcnt > n || rcnt > n || rcnt>lcnt {
			return
		}

		if lcnt == n && rcnt == n {
			ans = append(ans, s)
			return
		}

		dfs(lcnt+1,rcnt,s+"(")
		dfs(lcnt,rcnt+1,s+")")
	}

	dfs(0,0,"")

	return
}
// @lc code=end

