/*
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
 

示例 1：


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
示例 2：

输入：head = [], val = 1
输出：[]
示例 3：

输入：head = [7,7,7,7], val = 7
输出：[]
 

提示：

列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50

*/

/*
    思路：那个相等删那个
         相等:prev.next=head.next
         不相等:prev=head

    题解：递归
        边界:if(head==null) return head
         获取head.next=removeElements(head.next,val)
         head.val等于val：返回head.next
         不相等:返回head
        
*/

var removeElements = function(head, val) {
    let target=new ListNode()
    target.next=head
    let prev=target
    while(head){
        if(head.val==val){
            prev.next=head.next
        }
        else{
            prev=head
        }
        head=head.next

    }
    return target.next
};

var removeElements = function(head, val) {
    if(head==null){
        return head
    }
    head.next=removeElements(head.next,val)
    return head.val==val?head.next:head
};

/*
    思路：好久不递归，还挺不习惯的。简单的可以练习练习
*/