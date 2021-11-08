/*
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

 

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

/*
    思路：遍历链表，利用js的api，从头部进入数组。   

    题解：递归。理解递归本质上就是一种栈结构。
        先深入递归后再从尾到头返回val.

*/
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {number[]}
 */
 var reversePrint = function(head) {
    let target=[]
    while(head){
        target.unshift(head.val)
        head=head.next
    }
    return target
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

var reversePrint = function(head,arr=[]) {
    if(head){
        reversePrint(head.next,arr)
        arr.push(head.val)
    }
    return arr
}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：从栈想到递归特性，这种联动很赞呀。
*/
