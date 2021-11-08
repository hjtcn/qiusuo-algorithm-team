/*
剑指 Offer II 024. 反转链表
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。

 

示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]
 

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
 

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？



*/

/*
    思路：转着转着就绕里面了。
        踩坑：
         1.遍历跳出循环，不能拿head来计算。而应该拿更新后的head.next
         2.当前指针不能拿head，而是一个新的节点pointer去记录。
*/

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
 var reverseList = function(head) {
    if(!head) return head
    let pointer=head
    while(head.next){
        let next=head.next
        head.next=next.next
        next.next=pointer
        pointer=next
    }
    return pointer
};


var reverseList = function(head) {
    if(!head||!head.next) return head
    let next=reverseList(head.next)
    head.next.next=head
    head.next=null
    return next
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：一定要知道那个是不变的，借用变量辅助
                 那个是需要变的，同时可以跳出循环或递归。

*/