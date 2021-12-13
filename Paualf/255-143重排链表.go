定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 示例 1：
输入：head = [1,2,3,4]
输出：[1,4,2,3]

示例 2：
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：
链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000

1. Clarification:
在单链表的基础上，我没有想到更好的解决方法

想到了构造成双链表，为什么会想到双链表呢？因为从尾部节点要向前移动

2. 看题解：

使用线性表
将 节点放到 []*ListNode{} 里面
func reorderList(head *ListNode) {
    if head == nil {
        return
    }
    nodes := []*ListNode{}
    for node := head; node != nil; node = node.Next {
        nodes = append(nodes, node)
    }
    i, j := 0, len(nodes)-1
    for i < j {
        nodes[i].Next = nodes[j]
        i++
        if i == j {
            break
        }
        nodes[j].Next = nodes[i]
        j--
    }
    nodes[i].Next = nil
}

链表中点 + 链表逆序 + 合并链表

1. 找到原链表的中点
2. 将原链表的右半端反转
3. 将原链表的两端合并
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}

func mergeList(l1, l2 *ListNode) {
    var l1Tmp, l2Tmp *ListNode
    for l1 != nil && l2 != nil {
        l1Tmp = l1.Next
        l2Tmp = l2.Next

        l1.Next = l2
        l1 = l1Tmp

        l2.Next = l1
        l2 = l2Tmp
    }
}

func reorderList(head *ListNode) {
    if head == nil {
        return
    }
    mid := middleNode(head)
    l1 := head
    l2 := mid.Next
    mid.Next = nil
    l2 = reverseList(l2)
    mergeList(l1, l2)
}
3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 方法一给我的感觉是如果数组和链表放到一起，自己就有点搞不清楚了
4.2: 方法二给我的感觉是你怎么把一个问题给拆解分开，找到链表的中点 + 右半端反转 + 两端合并
