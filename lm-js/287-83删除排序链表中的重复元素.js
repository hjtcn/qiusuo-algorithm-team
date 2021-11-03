/*
83. 删除排序链表中的重复元素
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。

 

示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]
 

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列


*/

/*
    思路：由于是升序链表，因此只要对比相邻的两个节点(cur,next)就可以。while相同，则继续向后查询，直到不同，跳出循环后cur.next=next。
    同时，如果已经走到结尾，则直接跳出遍历链表的循环。
         如果没有走到结尾。继续更新cur和next。

    题解：看完题解后发现自己的想法复杂了。可以使用变量记录节点，但是要在确认需要的时候使用，否则用的太多反而把自己绕迷糊了。
         while(cur.next){
             if(cur.val==cur.next.val){
                 cur.next=cur.next.next
             }
             else{
                 cur=cur.next
             }
         }

*/
var deleteDuplicates = function(head) {
    if(!head) return head
    let cur=head,target=new ListNode()
    target.next=head
    next=cur.next
    while(next){
        while(next&&cur.val===next.val){
            next=next.next
        }
        cur.next=next
        if(!next) {
            break;
        }
        cur=next
        next=next.next
    }
    return target.next
};

var deleteDuplicates = function(head) {
    if(!head) return head
    let cur=head
    while(cur.next){
        if(cur.val==cur.next.val){
            cur.next=cur.next.next
        }
        else{
            cur=cur.next
        }
    }
    return head
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：在会的基础上简化优化。
*/