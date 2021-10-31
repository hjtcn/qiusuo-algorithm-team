实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。

1. Clarification:

1.1：快慢指针
没有想出来具体的解决方法

1.2: 遍历两次，第一次遍历找到总数，然后第二次遍历找到对应节点信息
遍历一次找出节点总数量
从头开始遍历到 n - k + 1 的节点，然后进行返回
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    total := getListTotal(head)

    for i := 1;i < total - k + 1;i++ {
        head = head.Next
    }

    return head.Val
}

func getListTotal(head *ListNode)(total int) {
    
    for head != nil {
        total++
        head = head.Next
    }

    return
}

2. 看题解：
快慢指针：
2.1: 创建虚拟头节点
2.2: 创建p,q两个指针，都指向虚拟头节点 dummyHead
2.3: 遍历k次，使p,q 结点间相距k个结点
2.4: 进行循环移动，每次都将p,q 向后一个结点移动
2.5: 当q指向nil 时，p所指结点即为所求结点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    dummy := &ListNode{}
    dummy.Next = head
    p,q := dummy,dummy

    for i := 0;i < k;i++ {
        q = q.Next
    }

    for q != nil {
        p = p.Next
        q = q.Next
    }

    return p.Val
}

这个解法更加的清晰
func kthToLast(head *ListNode, k int) int {
    slow, fast := head, head
    for k > 0 {
        fast = fast.Next
        k--
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    return slow.Val
}


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 思路知道但是没有具体的解决方法和不知道一样的，所以遇到问题，不要抱怨，尽可能的去给解决问题的方案和方法

