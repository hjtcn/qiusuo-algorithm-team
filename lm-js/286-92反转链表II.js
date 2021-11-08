/*
92. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]
 

提示：

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 

进阶： 你可以使用一趟扫描完成反转吗？

*/

/*
    思路：先掌握反转m~n链表，然后了解两个反转后的两个断点拼接
    踩坑：1.如果m==1，即从起始点就开始反转
        那么需要赋值 target.next=翻转后的头部节点
        2.遍历需反转的链表的条件用count++,是否>=m&&<n比较合适。
        我刚开始用了curNode.next的同时。如果>n就跳出循环(>n这个边界调节不够准确，也不太好理解)
          while(curNode.next){
            if(count<m||count>=n){
                break;
            }
            let next=curNode.next
            curNode.next=next.next
            next.next=pointer
            pointer=next
            count++;
         }

*/

var reverseBetween = function (head, m, n) {
    let target=new ListNode()
    let prev=new ListNode()
    target.next=head
    let count=1
    while(count<m){
        prev=head
        head=head.next
        count++
    }
    let tailLeft=prev
    let headRight=head
    let curNode=head
    let pointer=head
     while(count>=m&&count<n){
         let next=curNode.next
         curNode.next=next.next
         next.next=pointer
         pointer=next
         count++;
     }
     let headLeft=pointer 
     let tailRight=pointer.next  
     tailLeft.next=headLeft
     headLeft.next=tailRight
     if(m==1){
         //反转头部特殊处理
         target.next=headLeft
     }
     return target.next
 }


 /*
    时间复杂度：O(N)
    空间复杂度：O(N)
 */

/*
    思考：连着做的时候，发现对于反转链表掌控的更加熟练了，这道题就是在反转链表的基础上添加了断点拼接的细节处理。。

*/
