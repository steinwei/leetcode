/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package leetcode

// @lc code=start
func topKFrequent(nums []int, k int) (ans []int) {
    
	var swim func(a [][2]int, i int)
	swim = func(a [][2]int, i int) {
        if i == 0 {
            return
        }
        par := (i - 1) >> 1
        if a[par][1] > a[i][1] {
            a[par], a[i] = a[i], a[par]
            swim(a, par)
        }
	}

	var sink func(a [][2]int,i, heapSize int)
	sink = func(a [][2]int,i, heapSize int) {
        if i > heapSize {
            return
        }
        l, r := i << 1 + 1, i << 1 + 2
        if l < heapSize && a[l][1] < a[i][1] {
            a[l], a[i] = a[i], a[l]
            sink(a, l, heapSize)
        }

        if r < heapSize && a[r][1] < a[i][1] {
            a[r], a[i] = a[i], a[r]
            sink(a, r, heapSize)
        }
	}

    heap := [][2]int{}

    hashmap := map[int]int{}

    for _, v := range nums {
        hashmap[v] ++
    }

    for key, value := range hashmap {
        heap = append(heap, [2]int{key, value})
        heapSize := len(heap)
        swim(heap, heapSize - 1)
        if heapSize > k {
            heap[0] = heap[heapSize - 1]
            heap = heap[0: heapSize - 1]
            sink(heap, 0, len(heap))
        }
    }

    for _, v := range heap {
        ans = append(ans, v[0])
    }

    return 
}

// @lc code=end
