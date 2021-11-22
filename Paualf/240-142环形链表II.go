给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：
你是否可以使用 O(1) 空间解决此题？
 
示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

提示：
链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引

1. Clarfication:
找到环的入口点

快慢指针相遇的地方，然后让快指针，指向头指针，然后每次移动一步，相遇节点就是环的入口

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow,high := head,head
    flag := false

    for high != nil && high.Next != nil {
        slow = slow.Next
        high = high.Next.Next

        if slow == high {
            flag = true
            break
        }
    }

    // 如果快慢指针没有相遇，则代表没有环
    if !flag {
        return nil
    }

    high = head
    for slow != nil {
        if slow == high {
            return slow
        }
        slow = slow.Next
        high = high.Next
    }

    return nil
}
为什么这样可以呢？

从数学上解释下为什么呢？

自己画了下图，好像还是不太清楚。。。

2. 看题解：
看了下题解：
写了一遍公式，感觉模模糊糊吧

最后 a = (n-1) * (b+c) + c

感觉如果画图更容易理解这个地方


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 编程感觉还是逻辑，逻辑还是数学
