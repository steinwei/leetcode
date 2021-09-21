/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package leetcode

// @lc code=start
func leastInterval(tasks []byte, n int) int {

	n:= len(tasks)

	hashmap := map[byte]int{}
	kinds := 0
	maxb := 0
	

	for i, v := range tasks {
		if hashmap[v] == 0 {
			kinds++
		} else {
			
		}
		hashmap[v] ++
	}

	return 0
}

// @lc code=end
