/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */
package leetcode

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	var (
		dfs func(u int)
		// 0 未开始 1 遍历中 2 已完成
		visited = map[int]int{}
		// 判断是否有环
		valid = true
		edges = make([][]int, numCourses)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 1 {
				valid = false
				return
			} else if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}
		}	
		visited[u] = 2
		// 
		ans = append(ans, u)
	}

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		ans = []int{}
		return 
	}

	// reverse
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[numCourses-i-1] = ans[numCourses-i-1], ans[i]
	}

	return
}
// @lc code=end

