/*
 * @lc app=leetcode.cn id=77 lang=typescript
 *
 * [77] 组合
 */

// @lc code=start
function combine(n: number, k: number): number[][] {
    const res:Array<Array<number>>=[]

    const stack:Array<Array<number>>=[]
    
    let start:number = 1
    const end:number = n - k + 1
    
    if(k==1){
        while(start<=n){
            res.push([start++])
        }
        return res
    }else if(k<=0){
        res.push([])
        return res
    }

    while(start<=end){
        stack.push([start++])
    }

    while(stack.length!=0){
        const temp:Array<number> = stack.pop() ||[]
        let num:number=temp[temp?.length-1]
        for (let i = num+1; i <= n; i++) {
            const clone = temp.slice()
            clone.push(i)
            if(clone.length==k){
                res.push(clone)
            }else{
                stack.push(clone)
            }
        }

    }

    return res
};

// @lc code=end

