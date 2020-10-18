// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL


/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */
var reverseBetween = function(head, m, n) {
    //为什么要重新创建新节点？m是可以等于1的，如果一反转，你就根本找不到头节点了，所以要从原链表头节点的前一个节点开始记录。。。
    //一直没有想到。。。。
    let cur=node=new ListNode(),prev=null,w=1,tmp=head,lastNode=null
      cur.next=head//这里指创建的当前空节点的下一个节点为原链表的头节点
      //m之前的位置不用反转
      while(w<m){
           w++
           cur=cur.next
       }
       let front=cur//记录分割之前m位置节点的前一个节点-----不需要反转的最后一个点
        prev=tail=cur.next//记录m位置的节点----反转结束的最后一个点
        cur=prev.next
        //m～n之间的位置需要反转
        while(w>=m&&w<n){
            let kmp=cur.next
            cur.next=prev
            prev=cur
            cur=kmp
            w++
        }
        //开始拼接
        front.next=prev
        //我尝试打印了一下prev,会一直报错。Error - Found cycle in the ListNode，有一点点困惑，但是觉得是拼接的过程截取了环状以外剩余部分。       
        tail.next=cur
        return node.next
};

  /** 题解
分割三段。第一段和第三段不反转，记得存储连接处的节点，留着最后拼接，而第二段就直接反转即可
复杂度分析：
	时间复杂度：O(N)
    一次遍历，到n

	空间复杂度：O(n)
	变量存储了一些节点
 */