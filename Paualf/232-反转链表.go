给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]
 
提示：
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
 
进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

1. Clarification:
挨着一个一个的反转节点
还是需要注意节点的变化

1.1 记录next节点
1.2: 记录前一个节点
1.3: 将当前节点的next 赋予前一个节点
1.4: 当前节点后移

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    curr := head
    pre := nil

    for curr != nil {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }

    return pre
}

Line 10: Char 9: use of untyped nil (solution.go)

怎么声明空节点呢？

如果是递归怎么处理呢？
递归的话还是分解子问题：
使用数学归纳法，关注于眼前的问题

2. 看题解：
使用 varpre *ListNode 声明空节点就可以了

关键点：假设它已经反转
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：迭代：O(1)  递归：O(n)

4. 总结：
4.1: 写代码你的思维有时候需要解决当前的问题，然后再用数学归纳法去解决更大的问题，一开始不想那么复杂可能也是一件好事情
