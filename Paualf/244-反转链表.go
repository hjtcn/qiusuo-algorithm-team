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
需要一个前置节点 pre 记录
需要一个 curr 节点向后移动
需要一个next 节点指向curr.Next，然后赋值给curr向后移动

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pre,curr *ListNode
    curr = head

    for curr != nil {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }

    return pre
}

2. 看题解：

递归只思考当前，然后用数学归纳法去验证
想了下，也是，如果规模几十万，几百万的机器不可能都验证的，只有建立一个模型，然后证明这个模型是对的，然后推广
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：迭代O(1), 递归O(n)

4. 总结：
4.1: 多思考，多总结
