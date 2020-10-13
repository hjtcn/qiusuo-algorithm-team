// 给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

// 如果有两个中间结点，则返回第二个中间结点。

//  

// 示例 1：

// 输入：[1,2,3,4,5]
// 输出：此列表中的结点 3 (序列化形式：[3,4,5])
// 返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
// 注意，我们返回了一个 ListNode 类型的对象 ans，这样：
// ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
// 示例 2：

// 输入：[1,2,3,4,5,6]
// 输出：此列表中的结点 4 (序列化形式：[4,5,6])
// 由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
//  

// 提示：

// 给定链表的结点数介于 1 和 100 之间。


//暴力解决
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
var middleNode = function(head) {
    let two=head,num=0
    //遍历，先获取链表的长度
    while(head){
        head=head.next
        num++;
    }
    //获取一半长度，兼容长度为奇偶数的情况
    num=parseInt(num/2)
    while(num--){//指针跳转至中间位置
        two=two.next
    }
    return two
};

/** 题解
遍历。先遍历一遍，记录链表长度，再len/2获取中间节点的下标，再次从头指针开始遍历，获取中间节点
复杂度分析：
	时间复杂度：O(N)
    一层遍历

	空间复杂度：O(1)
	存取头指针和变量
 */

//快慢指针
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
var middleNode = function(head) {
    let slow=head,fast=head
    //循环条件保持fast&&fast.next，保证fast.next.next不会报错。
    while(fast&&fast.next){
        slow=slow.next
        fast=fast.next.next
    }
    return slow
};

/** 题解
快慢指针。等快指针遍历结束的时候，慢指针会正好走向中间节点
复杂度分析：
	时间复杂度：O(N)
    一层遍历

	空间复杂度：O(1)
	存取头指针
 */