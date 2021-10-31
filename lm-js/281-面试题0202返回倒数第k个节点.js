/*
面试题 02.02. 返回倒数第 k 个节点
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。
*/

/*
    思路:滑动窗口，保证窗口长度为k,当右指针到链表尾部时，左指针就是目标节点啦。
*/
var kthToLast = function(head, k) {
    let slow=head,fast=head
    while(k--){
        fast=fast.next
    }
    while(fast){
        slow=slow.next
        fast=fast.next
    }
    return slow.val
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：简洁的代码看起来就很开心。链表尽量避免用数组去处理题目
*/