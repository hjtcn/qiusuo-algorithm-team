给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
示例 2：
输入：head = [], val = 1
输出：[]
示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]
 
提示：
列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50

1. Clarification:
    和昨天的题蛮像的，不一样的地方就是昨天是删除所有重复的元素，这道题是要删除所有值等于val的元素
    解法感觉昨天的题可以复用
    目标是：
        如果元素的值 == val，就要把这个节点给删除
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode {
        Next:head,
    }
    cur := dummy
    for cur != nil {
          if cur.Next != nil && cur.Next.Val == val {
            for cur.Next != nil && cur.Next.Val == val {
                cur.Next = cur.Next.Next
            }
        }else {
            cur = cur.Next
        }
    }
  
    return dummy.Next
}

上面的代码一开始漏掉了 for cur != nil 

2. 看题解：
题解中的递归也是一种思路，值得参考下
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    head.Next = removeElements(head.Next, val)
    if head.Val == val {
        return head.Next
    }
    return head
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度： 递归：O(n) 迭代 ：O(1)

4. 总结：
4.1: 一开始以为链表删除元素挺简单的，做完以后才发现是自己挺简单的，所以不要个人觉得，去尝试，去享受反馈，不抵抗，享受过程

4.2: 每个人都有自己的课题和知识点要过
