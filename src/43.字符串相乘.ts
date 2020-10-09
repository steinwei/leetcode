/*
 * @lc app=leetcode.cn id=43 lang=typescript
 *
 * [43] 字符串相乘
 */

// @lc code=start
function multiply(num1: string, num2: string): string {
    let res = ''

    let len1 = num1.length - 1, len2 = num2.length - 1
    
    // 每个数循环
    let flag = 0
    while(len1>= 0 || len2 >= 0){
        // end
        let ans = 0
        if(len1>=0&&len2>=0){
            ans = +num1[len1] * +num2[len2] + flag
        }else if(len1>=0){
            ans = +num1[len1] + flag
        }else{
            ans = +num2[len2] + flag
        }

        flag = Math.floor(ans / 10)

        // res
        res = ans % 10 + res

        // loop
        len1--
        len2--

    }


    return res
};
// @lc code=end

// multiply("123456789", "987654321")