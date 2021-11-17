/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package leetcode

// @lc code=start
func maxSlidingWindow(nums []int, k int) (ans []int) {
	var (
		// 单调递减队列
		p = make([]int, k)
		push func(index int)
		pop func(val int)
		n = len(nums)
	)

	push = func(index int) {
		for len(p) > 0 && nums[index] > nums[p[len(p)-1]] {
			p = p[:len(p)-1]
		}

		p = append(p, index)
	}

	pop = func(index int) {
		for len(p) > 0 && p[0] == index {
			p = p[1:]
		}
	}

	for i := 0; i < k - 1; i++ {
		push(i)
	}

	for i := k - 1; i < n; i++ {
		push(i)
		ans = append(ans, nums[p[0]])
		pop(i-k+1)
	}
	
	return
}
// @lc code=end

