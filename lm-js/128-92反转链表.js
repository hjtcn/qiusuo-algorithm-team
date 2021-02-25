// 92. 反转链表 II
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL



/** 
 思路：这个题思路都清晰了了，主要就是细节问题

    首先指定l--->r的链表。也就是l前一个节点记住，r后一个节点记住，用来拼接。
    需反转链表反转即可。

    该题还需考虑的一个细节点就是如果从第一个节点就开始反转的情况。

    注意细节
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
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
var reverseBetween = function(head, left, right) {
    let prev=front=new ListNode()
    prev.next=head
    let count=1,flag=1
    while(count<left){
        count++;
         front=head
         head=head.next
    }
    let  last=head,next=head.next
    while(count>=left&&count<right){
        count++;
        let last=head
        head=next
        next=head.next
        head.next=last
    }
    if(left==1){
        prev.next=head
    }
    front.next=head
    last.next=next
    return prev.next
};

/*
时间复杂度：O(N)
一次链表遍历
空间复杂度:O(N)
声明变量记录节点
*/