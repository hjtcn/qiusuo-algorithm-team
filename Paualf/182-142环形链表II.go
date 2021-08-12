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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

    1. Clarification:
        思路是双指针，但是怎么用呢？
        纸上画了下：
            让快指针，每次走2步，慢指针每次走一步，两者相交的地方时候，再让慢指针从头开始走，慢指针和快指针每次只移动一步，两者相交则是环的入口
        我目前推的是快指针比慢指针多走一个环的长度，如果有环的话
        感觉证明的话还是要来一下数学公式才行

func detectCycle(head *ListNode) *ListNode { 
    if head == nil || head.Next == nil {
        return nil
    }
    slow,fast := head,head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    slow = head
    for slow != nil && fast != nil {
        if slow == fast {
            return slow 
        }
        slow = slow.Next
        fast = fast.Next
    }
    return nil
}

2. 看题解：
hash 表还是蛮厉害的，没有想到，只想到了快慢指针，没有把所有思路想一下，漏了一些
官方题解的图和数学公式都很到位
最后推导出来
a  = c + （n-1)(b+c) ，其中 (n-1)(b+c) 是环长，可以去掉，所有a = c ，所有从快慢指针相遇点，再进行移动再相遇就是它们的交点了

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 数学好像还是蛮重要的，丢了好久了。。。，有需要的话慢慢补补数学
