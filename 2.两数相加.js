/*
 * @lc app=leetcode.cn id=2 lang=javascript
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    function generateNode(arr){
        function listNode(data={},next=null){
            return {
                data:data,
                next:next
            }
        }
        
    }
    function parser(node,list){
        var data = node.data;
        var next = node.next;
        var list = list||[];
        list.push(data);
        if(!next) return list;
        parser(next,list);
    }
    function transform(arr){
        var arr1= ([arr[0],arr[1],arr[2]]=[arr[2],arr[1],arr[0]]);
        return parserInt(arr1.toString().replace(/,/g,''));
    }
    function distract(){
        
    }
    // pipeline
    function main(l1,l2){
        var list1=new parser(l1);
        var list2=new parser(l2);
        
    }
    return main(l1,l2);
};
// @lc code=end

function listNode(data={},next=null){
    return {
        data: data,
        next: next
    }
}

var l1 = new listNode(2,new listNode(4,new listNode(3)));
var l2 = new listNode(5,new listNode(6,new listNode(4)));

addTwoNumbers(l1,l2);

// console.log()