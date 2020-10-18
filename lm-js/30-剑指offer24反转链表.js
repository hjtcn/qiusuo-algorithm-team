// 剑指 Offer 24. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

 

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
 

// 限制：

// 0 <= 节点个数 <= 5000


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
var reverseList = function(head) {
    let cur=head,prev=null,tmp
     while(cur){
         //借用变量存储下一个节点，用来保证跳转，运转while循环
         tmp=cur.next
         //开始反转节点。当前节点的next节点为上一个节点
        cur.next=prev
        //记录前一个节点，下一次循环使用
        prev=cur
        //当前节点指向下一个节点，进行while循环
        cur=tmp
     }
     //prev当前是指向最后一个节点，而cur已经指向null
     return prev
};

/** 题解
迭代，借用变量改变指向
复杂度分析：
	时间复杂度：O(N)
    一次遍历

	空间复杂度：O(1)
	存取指针
 */


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
var reverseList = function(head) {
    // 递归结束
    if (head === null || head.next === null) {
      return head
    }
  
    // 递归反转 子链表
    let newReverseList = reverseList(head.next)
    // 获取原链表的第2个节点newReverseListTail
    let newReverseListTail = head.next
    // 调整原来头结点和第2个节点的指向
    newReverseListTail.next = head
    head.next = null
  
    // 将调整后的链表返回
    return newReverseList
  };

  /** 题解
递归.只关注前两个节点，进行翻转，直至链表结束。
复杂度分析：
	时间复杂度：O(N)
    一次遍历，深度优先直至链表结束。

	空间复杂度：O(n)
	变量存储了一条新的链表
 */


 //翻题解扒到了一个满秀的解构赋值的写法。

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
var reverseList = function(head) {
    //也是借助变量。
    [curr, pre] = [head, null];
    //结构赋值直接进行翻转
    while(curr) [curr.next, pre, curr] = [pre, curr, curr.next];
    return pre;
};

/** 题解
迭代，借用变量改变指向
复杂度分析：
	时间复杂度：O(N)
    一层遍历

	空间复杂度：O(1)
	存取指针
 */