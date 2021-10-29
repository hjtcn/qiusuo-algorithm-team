/*
    剑指 Offer II 077. 链表排序
给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

 

示例 1：



输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：



输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]
 

提示：

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
 

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？


*/

/*
    思考：想到了借用数组
        也看到了进阶方案。时间复杂度为O(NLogN).
        有哪些排序方法的时间复杂度为O(NLogN)
        1.快排
        2.归并排序
        没写出来方案，看了题解。
        1.归并排序。先拆后合。
        拆：快慢指针。找到中间节点midNode。怎么拆？midNode.next=null  
        合:有序链表合并。
        使用递归。从底向上。从两个节点拼接，到两条有序链表拼接。
        踩坑：1.借用空节点当头部节点。
             2.断开链表
             3.先从合并两个有序链表考虑，还要考虑最后剩余的有序链表拼接
*/
let getMidNode=(root)=>{
    let slow=root,fast=root,prev=root
    while(fast&&fast.next){
        prev=slow
        slow=slow.next
        fast=fast.next.next
    }
    //截断，还要借用额外变量，防止slow被修改。
    prev.next=null
    return slow
}
let merge=(l,r)=>{
    let node=new ListNode(-1)
    let curNode=node
    while(l&&r){
        if(l.val>r.val){
            curNode.next=r
            r=r.next
        }
        else{
            curNode.next=l
            l=l.next
        }
        curNode=curNode.next
    }
    //这个地方也很重要，l和r的链表长度不相等，最后剩一个拼接上。
    if(l) curNode.next=l
    if(r) curNode.next=r
    return node.next
}
var sortList = function(head) {
    if(!head||!head.next) return head
    let midNode=getMidNode(head)
    let left=sortList(head)
    let right=sortList(midNode)
    return merge(left,right)
}
/*
    思考：多扣细节，排序真的蛮重要的。

*/