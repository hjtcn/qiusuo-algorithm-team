给定一个链表，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。
进阶：
你能用 O(1)（即，常量）内存解决此问题吗？

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
 
提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

1. Clarfication:

链表是否有环：
快慢指针，快指针，慢指针
我有的疑问是为什么快指针赶上慢指针就有环了呢？

slow,high := head

for hight != nil {
     slow = slow.Next
high = hight.Next.Next
if slow == high {
return true
}
}

return -1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
     slow,high := head,head

    for high != nil && high.Next != nil {
        slow = slow.Next
        high = high.Next.Next
        if slow == high {
            return true
        }
    }

    return false
}

2. 看题解：
没想到使用map使用
func hasCycle(head *ListNode) bool {
    seen := map[*ListNode]struct{}{}
    for head != nil {
        if _, ok := seen[head]; ok {
            return true
        }
        seen[head] = struct{}{}
        head = head.Next
    }
    return false
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 现在感觉把技术方案写的清晰并且能让人看懂，感觉比写代码难
4.2: 现在知道使用快慢指针解决问题了，但是为什么要使用快慢指针来解决呢？感觉要到数学问题上看了
