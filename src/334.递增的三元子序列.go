/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */
package leetcode

// @lc code=start
func increasingTriplet(nums []int) bool {
	// var (
	// 	n = len(nums)
	// 	first, second = nums[0], math.MaxInt32
	// )

	// if n < 3 {
	// 	return false
	// }

	// // greedy
	// // core: if third < second > first, then second = tihrd [update]
	// // if third > second > first then true
	// // if both wrong, then reset first
	// for i := 1; i < n; i++ {
	// 	curr := nums[i]
	// 	if curr > second {
	// 		return true
	// 	} else if curr > first {
	// 		second = curr
	// 	} else {
	// 		first = curr
	// 	}
	// }

	// return false

	// second solutio
	// make sure that there is solution what improve min<cur<max become true
	var (
		n = len(nums)
		leftminn = make([]int,n)
		rightmaxn = make([]int,n)
		min func(x,y int) int
		max func(x,y int) int
	)

	min = func(x, y int) int {
		if x>y {
			return y
		}
		return x
	}

	max = func(x, y int) int {
		if x>y {
			return x
		}
		return y
	}

	leftminn[0] = nums[0]
	for i := 1; i < n; i++ {
		leftminn[i] = min(leftminn[i-1], nums[i])
	}

	rightmaxn[n-1] = nums[n-1]
	for i := n-2; i > -1; i-- {
		rightmaxn[i] = max(rightmaxn[i+1], nums[i])
	}

	for i := 0; i < n; i++ {
		if nums[i]>leftminn[i]&&nums[i]<rightmaxn[i] {
			return true
		}
	}

	return false
}
// @lc code=end

