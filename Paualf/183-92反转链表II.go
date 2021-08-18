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
想了下没想出来

2. 看题解：
看了题解以后知道思路了，但是让我自己写真的能写出来么，不好说
里面细节还是蛮多的，思路都没有理解清楚可能

第一种方法：
思路：
记录左节点
记录右节点
切出中间一个链表，对中间链表进行反转
然后将链表连接到一起

第2种方法：
    使用pre,curr,next 记录反转过程：
curr : 指向待反转区域的第一个节点left
next: 永远指向curr 的下一个节点
pre: 永远指向待反转区域的第一个节点left的前一个节点
说实话第2种方法下一次遇到还不一定会写
还有链表的反转都没有写出来。。。

func reverseLinkedList(head *ListNode) {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
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
4.1: 链表里面细节挺多的，又想了下，真的是细节么，还是真的就是自己不会。。。 应该是真的是自己不

4.2: 单纯反转链表，可以注意三者的变换，pre,curr,next 的变换=
