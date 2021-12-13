给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 
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

1.  Clarification:
反转链表，一开始的时候没有想清楚，纸上画了一下才清楚了，想不清楚的时候纸上画一下，思路清晰了一下

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode

    cur := head

    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    
    return pre
}

2. 看题解：
递归是反向工作，假设链表的其余部分已经被反转
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
感觉还是数学归纳法的思想，假设后面好了，然后推理论证

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 想不清楚的画一下图去理解
4.2: 解决问题的时候不要想太多，一步步解决就好了
4.3: 尊重自己，尊重他人
