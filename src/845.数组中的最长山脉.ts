/*
 * @lc app=leetcode.cn id=845 lang=typescript
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
function longestMountain(A: number[]): number {
    let prev = A[0], track = 1, ret = 0, peek = 0, direction = 0

    // 理解错题意， 以为必须对称
    for (let i = 1; i < A.length; i++) {
        const last = A[i];
        // peek
        console.log(prev, last, track, direction, peek)
        if(last > prev) {
            if(direction == -1) {
                if(peek) {
                    ret = Math.max(ret, track)
                }
                track = 1
            }

            if(last > A[i + 1]) {
                peek = 1
            }

            track ++
            direction = 1
            prev = last
        } else if(last < prev){
            track ++
            direction = -1
            prev = last
            // ret
            if(i == A.length - 1 && peek) {
                ret = Math.max(ret, track)
            }
        } else {
            if(direction == -1) {
                if(peek) {
                    ret = Math.max(ret, track)
                }
            }
            track = 1
        }
    }

    return ret < 3 ? 0 : ret
};
// @lc code=end

