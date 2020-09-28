// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

//  

// 示例：

// 给定一个链表: 1->2->3->4->5, 和 k = 2.

// 返回链表 4->5.


/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
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
 * @param {number} k
 * @return {ListNode}
 */
var getKthFromEnd = function(head, k) {
    //定义两个变量，count用来维持链表长度，tag用来记录当前指针。
    let count=0,tag=head;
    //当当前指针指向null,即原链表已经走完
    while(tag){
        if(count<k){//使链表的长度维持到k
            count++
        }
        else{//当链表的长度到k，则头指针也每次循环后移一位
            head=head.next
        }
        tag=tag.next//当前指针不断后移，直至最后一位
    }
    return head//倒数第k个节点就为改变后的头指针
};

/**题解
   维持链表长度为k，少于k,仅将当前指针后移，大于等于k,则把后移当前指针的同时，头指针也要后移，直至链表走完。
 
  复杂度分析：
    时间复杂度是:O(N)
    一层循环将该链表遍历完毕。
    空间复杂度：O(1）
    定义两个变量，存放链表长度，以及记录当前指针。
 */
