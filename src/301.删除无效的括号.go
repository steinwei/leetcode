/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package leetcode

// @lc code=start
func removeInvalidParentheses(s string) (ans []string) {
	var (
		n = len(s)
		dfs func(start, score int, tmp string)
		lcount = 0
		rcount = 0
		maxLen = 0
		set = map[string]struct{}{}
	)

	for _, v := range s {
		if v == '(' {
			lcount ++
		}

		if v == ')' {
			rcount ++
		}
	}

	max := min(lcount, rcount)
	
	dfs = func(start, score int, tmp string) {
		if score > max || score < 0 {
			return
		}
		if start == n {
			if score == 0 && len(tmp) >= maxLen {
				if len(tmp) > maxLen {
					for k := range set {
						delete(set, k)
					}
				}
				maxLen = len(tmp)
				set[tmp] = struct{}{}
			}
			return
		}

		x := string(s[start])
		if x == "(" {
			dfs(start+1, score+1, tmp + x)
			dfs(start+1, score, tmp)
		} else if x == ")" {
			dfs(start+1, score-1, tmp + x)
			dfs(start+1, score, tmp)
		} else {
			dfs(start+1, score, tmp + x)
		}			
	}

	dfs(0, 0, "")

	for k := range set {
		ans = append(ans, k)
	}

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
// @lc code=end

