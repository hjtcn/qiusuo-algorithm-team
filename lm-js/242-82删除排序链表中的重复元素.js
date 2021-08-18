/*
82. 删除排序链表中的重复元素 II
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。

返回同样按升序排列的结果链表。

 

示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]
 

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列
*/

/*
    思路：记录三个节点prevNode,head,nextNode
         进行比较是否是重复值。
         同时需要prev节点：
         记录当前head节点若重复，则prev.next=null
         若不重复，则prev.next=head,prev=head
         用prev将新的链表串起来
         坑点：1.循环的条件为while(head&&head.next)
            会导致最后一个节点未处理，因此，跳出循环后，比较head和prevNode是否相等
            不相等:prev.next=head
            相等：不处理，prev没有改变
            2.链表的节点值可能为0，我初始化prev时，比较节点值可能因为head.val=0，而导致虚拟头节点和head.val相等，head被删掉。
            因此我选择了初始化prev的值为101.大于题目的最大节点值
    题解：一个while循环控制节点迭代，一个while循环比较节点重复。一旦节点重复，就删除所有重复的节点
        但是比较巧妙的是，比较的是cur.next和cur.next.next
        cur不仅是当前值，还能一直保持不重复

*/

var deleteDuplicates = function(head) {
    if(!head){
        return head
    }
    let target=prev=new ListNode(101)
    target.next=head
    prev.next=head
    let nextNode=head.next
    let prevNode=prev
    while(head&&nextNode){
        if(head.val==prevNode.val||head.val==nextNode.val){
            prev.next=null
        }
        else{
            prev.next=head
            prev=head
        }
        prevNode=head
        head=head.next
        nextNode=head.next
    }
    if(head.val!=prevNode.val){
        prev.next=head
    }
    return target.next
};


var deleteDuplicates = function(head) {
    let target=new ListNode()
    target.next=head
    let cur=target
    while(cur.next&&cur.next.next){
        if(cur.next.val==cur.next.next.val){
            let x=cur.next.val
            while(cur.next&&cur.next.val==x){
                cur.next=cur.next.next
            }
        }
        else{
            cur=cur.next
        }
    }
    return target.next
};

/*
    思考：我刚开始思考的时候也是有两层while循环的，但是搞着搞着就晕了，最后又把最外层的剥离掉，重新走拼接思路。
        太多细节要控制了

*/