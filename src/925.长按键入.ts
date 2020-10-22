/*
 * @lc app=leetcode.cn id=925 lang=typescript
 *
 * [925] 长按键入
 */

// @lc code=start
function isLongPressedName(name: string, typed: string): boolean {
    let first = 0
    let second = 0

    while(second < typed.length) {
        // console.log(name[first], typed[second])

        if(name[first] == typed[second]) {
            first ++
            second ++
        }  else if(typed[second] == typed[second - 1]){
            second ++
        } else {
            break
        }
    }

    return first == name.length && second == typed.length;
};
// @lc code=end

