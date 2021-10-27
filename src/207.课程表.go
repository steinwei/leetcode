/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
package leetcode

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid = true
		dfs func(u int)
	)

	dfs = func(u int) {
		visited[u]= 1

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
	}

	// 建图
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	// 遍历图
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}
// @lc code=end

