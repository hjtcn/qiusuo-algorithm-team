/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

 

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
 

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

*/

/*
    思路：两次扫描。特殊处理如果n==size的情况。其余直接赋值node.next=node.next.next即可

        一次扫描。感觉更像滑动窗口多一些，在遍历过程中，count++,当count=n时，窗口起始位置就是要被删除的节点了，count开始>n，一直更新窗口的起始位置,及它的前一节点prev，作为要被删除的节点targetDel。
        跳出循环，prev.next=targetDel.next
        依然要特殊处理需要删除头节点的情况。
        1.删除头节点。返回head.next
        2.删除非头节点。返回head
        此时踩了一个小坑。如何判断头节点??
        还是需要根据size和n进行比较比较确定。反正已经遍历一遍了，可以确定size了
        第一次想的单纯比较head.val和targetDel.val，后来发现，链表过程中会出现相同的val.
        



*/

/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
 var removeNthFromEnd = function(head, n) {
    let target=head
    let cur=head,size=0,count=0,node=head
    while(cur){
        cur=cur.next
        size++;
    }
    if(n==size){
        target=head.next
        return target
    }
    while(node){
        let prev=node
        node=node.next
        count++
        if(count==size-n){
            prev.next=node.next
        }
    }
    return target
};


var removeNthFromEnd = function(head, n) {
    let target=head,targetDel=head
    let cur=head,count=0,node=head,prev=null
    while(cur){
        cur=cur.next
        count++;
        if(count>=n){
           prev=targetDel
            targetDel=node
            node=node.next
        }
    }
    prev.next=targetDel.next
    if(count==n){
        return target.next
    }
    else{
        return target
    }

};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：链表记录前一节点，当前节点，下一节点是非常重要的。我想到了动规的降维，用的比之前顺畅了
*/