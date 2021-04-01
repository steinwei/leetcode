/*
 * @lc app=leetcode.cn id=844 lang=typescript
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
function backspaceCompare(S: string, T: string): boolean {
    let l1 = S.length - 1
    let l2 = T.length - 1

    let c1 = 0
    let c2 = 0

    let s1: string = ''
    let s2: string = ''

    for (let i = l1; i > -1; i--) {
        const e = S[i]
        if(e == '#') {
            c1 ++
        } else if(c1){
            c1 --
        } else {
            s1 = e + s1
        }
    }
    
    for (let i = l2; i > -1; i--) {
        const e = T[i]
        if(e == '#') {
            c2 ++
        } else if(c2){
            c2 --
        } else {
            s2 = e + s2
        }
    }

    return s1 == s2
};
// @lc code=end

