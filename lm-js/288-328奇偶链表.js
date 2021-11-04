/*
328. 奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL 
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。   

*/

/*
    思路：1.重新更新奇数节点的next节点拼接
         2.最后完整奇数链表拼接偶数链表

         踩坑：遍历一次出来，你知道该链表最后的节点是奇数节点还是偶数节点？？如何保证奇数在前，偶数在后呢？
    题解: 官方题解中，提前就分成两条链表了，处理起来有点不是那么得心应手。
*/

/*
var oddEvenList = function(head) {
    if(!head) return head
    let target=head
    let evenHead=target.next
    let prev=null,count=1
    while(head.next){
        let next=head.next
        head.next=next.next
        prev=head
        head=next
        count++;
    }
    if(count%2){
        head.next=evenHead
    }
    else{
        prev.next=evenHead
    }

    return target
};

*/

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：有时候一分割，细节上就有些控制不住了。
*/