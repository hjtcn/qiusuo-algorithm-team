/*
92. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]
 

提示：

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 

进阶： 你可以使用一趟扫描完成反转吗？
*/
/*
    思路：
    1.记录连接点。
        left的前一个节点(frontLink)
        left节点(nextLink)
        right节点(cur)
        right的后一个节点(next）
        frontLink.next=cur
        nextLink.next=next
    2.left和right之间节点进行反转，直到cur指向原始right指向的位置
        let next=cur.next
        while(count<=left&&count>right){
            //内部反转
            let prev=cur
            cur=next  
            next=cur.next  
            cur.next=prev
            count
        }
    题解：一次扫描。穿针引线。
        记录三个参数
        1.cur,待反转区的第一个节点
        2.prev,left的前一个节点，永远不变
        3.next,cur的下一个节点，cur变，next跟着变
        在待反转区[left,right)有三个操作
        let next=cur.next
        1.cur.next=next.next
        2.next.next=prev.next
        3.prev.next=next
        有点巧妙，处理上比前一种拼接且不断反转要优多了
        其实仔细看穿针引线的三个操作。
        如1 2 3 4 中，2和3进行反转
        先2的next指到4 拼后
        再3的next指到2 开始核心两个节点反转
        再1的next指到3 接前

*/
/**
 * @param {ListNode} head
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
 var reverseBetween = function (head, m, n) {
    let prevTarget = new ListNode()
    prevTarget.next = head
    let cur = head, prev = prevTarget
    let count = 1, frontLink = new ListNode()
    while (count < m) {
        prev = cur
        cur = cur.next
        count++;
    }
    frontLink = prev
    let next = cur.next, nextLink = cur
    while (count >= m && count < n) {
        let tail = cur
        cur = next
        next = cur.next
        cur.next = tail
        count++;
    }
    frontLink.next = cur
    nextLink.next = next
    return prevTarget.next
}

var reverseBetween = function (head, m, n) {
    let prevTarget = new ListNode()
    prevTarget.next = head
    let cur = head, prev = prevTarget
    let count = 1
    while (count < m) {
        prev = cur
        cur = cur.next
        count++;
    }
    while (count >= m && count < n) {
        let next=cur.next
        cur.next=next.next
        next.next=prev.next
        prev.next=next
        count++;
    }
    return prevTarget.next
}
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/
/*
    思考：链接一定要注意处理细节。比如记录prev节点，next节点，反转过程中需要拼接的节点。
*/