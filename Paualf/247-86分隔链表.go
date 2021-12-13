给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

 示例 1：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：
输入：head = [2,1], x = 2
输出：[1,2]
 

提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200

1. Clarification:
没有啥思路


2. 看题解：
维护两个链表 small和large,small 链表按照顺序存储所有小于x的节点，large链表按顺序存储所有大于等于x的节点，遍历完原链表后，将small 链表尾节点指向large 链表头节点

func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    smallHead := small
    large := &ListNode{}
    largeHead := large
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        } else {
            large.Next = head
            large = large.Next
        }
        head = head.Next
    }
    large.Next = nil
    small.Next = largeHead.Next
    return smallHead.Next
}
特别容易出错的地方：   large.Next = nil，一不小就 panic 了

3. 复杂度分析
时间复杂度：O(n)
空间复杂度：O(1)


4. 总结：
4.1: 感觉链表对分析要求比较高，不小心就运行失败了
