给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]
 
提示：
链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 
进阶： 你可以使用一趟扫描完成反转吗？

1. Clarification:
这种反转链表，想到了穿针引线的方式，自己手动在纸上画的时候，发现只是知道有这个东西，实现不出来。。。

2. 看题解：
关注题解中图中 pre  left right curr 节点的地方

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }
}


func reverseBetween(head *ListNode,left,right int)*ListNode {
	 // 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head

    pre := dummyNode
    // 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
    // 建议写在 for 循环里，语义清晰
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }

    // 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
    rightNode := pre
    for i := 0; i < right-left+1; i++ {
        rightNode = rightNode.Next
    }

    // 第 3 步：切断出一个子链表（截取链表）
    leftNode := pre.Next
    curr := rightNode.Next

    // 注意：切断链接
    pre.Next = nil
    rightNode.Next = nil

    // 第 4 步：同第 206 题，反转链表的子区间
    reverseLinkedList(leftNode)

    // 第 5 步：接回到原来的链表中
    pre.Next = rightNode
    leftNode.Next = curr
    return dummyNode.Next
}
穿针引线：头插法
curr: 指向待反转区域的第一个节点left;
next: 永远指向curr的下一个节点，循环过程中，curr 变化以后next会变化
pre: 永远指向待反转区域的第一个节点left的前一个节点

关键点：pre 不动，curr 通过变化向后移动
func reverseBetween(head *ListNode, left, right int) *ListNode {
    // 设置 dummyNode 是这一类问题的一般做法
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head
    pre := dummyNode
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
    for i := 0; i < right-left; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return dummyNode.Next
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 没有容易得到的东西，耐心一点，诚恳一点，用心把事情做好
