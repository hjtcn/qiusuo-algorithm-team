// 82. 删除排序链表中的重复元素 II
// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

// 示例 1:

// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
// 示例 2:

// 输入: 1->1->1->2->3
// 输出: 2->3

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function (head) {
    //利用map对象进行value与出现次数的存储
    let cur=head,m=new Map(),again=head,list=new ListNode(),res=list
    //遍历一次链表，进行存储
    while(cur){
     m.set(cur.val,m.has(cur.val)?m.get(cur.val)+1:1)
     cur=cur.next
    }
    //遍历赋值出现次数只为1的value的节点
    while(again){
        if(m.get(again.val)==1){
            //new 一个目标value的新对象作为新的节点
         list.next=new ListNode(again.val)
         list=list.next
        }
        again=again.next
    }
    //提前存储目标链表头节点，并从头节点的next节点开始返回。
    //因为again节点现在已经指向最后的尾节点了
    return res.next
 };

   /** 题解
   利用map对象进行存储查询，创建新链表仅存储value次数为1的。
   复杂度分析：
	时间复杂度：O(N)
    两次遍历，一次存储，一次查询

	空间复杂度：O(n)
	变量声明了map对象。
 */