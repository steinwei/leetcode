/*
 * @lc app=leetcode.cn id=231 lang=typescript
 *
 * [231] 2的幂
 */

// @lc code=start
function isPowerOfTwo(n: number): boolean {
    if ( n <= 0) return false
    let res = false
    let e = n
    while(true) {
        if( e == 1){
            res = true
            break
        } else if (e < 1) {
            break
        }
        e = e / 2
    }
    return res
};
// @lc code=end

