存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。
示例 1：
输入：head = [1,1,2]
输出：[1,2]

示例 2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

1. Clarification:
删除排序链表中的重复元素，排序的想到了指针的移动

一个节点指向当前节点
一个节点指向当前节点后的一个节点
如果next 节点 == 当前节点，则移动next 指针，不想等的时候更改 curr的Next,然后current 指针后移

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
     curr := head
     for curr != nil {
         next := curr.Next
         for next != nil && curr.Val == next.Val {
             next = next.Next
         }
         curr.Next = next
         curr = curr.Next
     }
     return head
}

2. 看题解：

题解判断的是 cur.Next != nil 也有一定的道理的
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := head
    for cur.Next != nil {
        if cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }

    return head
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

总结：
4.1: 多思考，多画图，不要一上来就操作，那样往往会无功而返
