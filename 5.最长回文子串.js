/*
 * @lc app=leetcode.cn id=5 lang=javascript
 *
 * [5] 最长回文子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    // 翻转方法
    function reverse(string){
        var _s = '';
        var len = string.length;
        for(var i = len - 1; i>= 0;i--){
            _s=string[i]
        }
        return _s
    }
    // 遍历
    function traversal(){
        
    }

    function equals(a,b){
        return a === b;
    }
    // main
    function longest(a){
        
    }
    // 用来缓存长度 丢弃多余的长度
    function cache(max){
        this.max = max;
    }

    function getMax(){
        return this.max;
    }

    return longest(s);
};
// @lc code=end

