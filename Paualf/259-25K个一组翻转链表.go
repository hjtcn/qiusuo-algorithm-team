给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

示例 3：
输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]

示例 4：
输入：head = [1], k = 1
输出：[1]

提示：
列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz

1. Clarification:
翻转链表

题目不难，还是不会写。。。

2. 看题解:
题目不难，写好不简单
func reverseKGroup(head *ListNode, k int) *ListNode {
    hair := &ListNode{Next: head}
    pre := hair

    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if tail == nil {
                return hair.Next
            }
        }
        nex := tail.Next
        head, tail = myReverse(head, tail)
        pre.Next = head
        tail.Next = nex
        pre = tail
        head = tail.Next
    }
    return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
    prev := tail.Next
    p := head
    for prev != tail {
        nex := p.Next
        p.Next = prev
        prev = p
        p = nex
    }
    return tail, head
}

反转链表的时候和正常的加头节点的还有一点不一样，因为判断终止的条件不一样了

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 多画图，想不清楚的话就多画下图看下


