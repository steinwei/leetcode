/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package leetcode

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	var (
		n = len(tasks)
		visited = make(map[byte]int)
		ans = 0
	)

	for _, v := range tasks {
		if visited[byte(v)] == 0 {
			
		} else {
			visited[byte(v)] = 1
		}
		visited[byte(v)] = 2
		ans ++
	}

	return ans
}

// @lc code=end
