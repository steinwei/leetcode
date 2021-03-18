/*
 * @lc app=leetcode.cn id=1002 lang=typescript
 *
 * [1002] 查找常用字符
 */

// @lc code=start
function commonChars(A: string[]): string[] {
    let res: string[] = []
    
    
    while(A.length) {
        const  next = A.pop() || ''
        res = !res.length ? next.split('') : next.split('').filter((s:string)=>{
            const exist = res.indexOf(s)
            if(exist > -1) {
                return true
            }
            return false
        })

    }

    return res
};
// @lc code=end


// ["acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"]
// ["cool","lock","cook"]
// ["bbddabab","cbcddbdd","bbcadcab","dabcacad","cddcacbc","ccbdbcba","cbddaccc","accdcdbb"]
