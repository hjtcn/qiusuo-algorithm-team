给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]
提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109
1. Clarification:
思路：
找到最后一个节点，将最后一个节点指向第一个节点
如果只有一个节点，直接返回
旋转k次，执行k次旋转即可
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    for i := 0;i < k;i++ {
        head = rotate(head)
    }
    return head
}
func rotate(head *ListNode) *ListNode{
    // 找到最后一个节点的前一个节点断开
    cur := head
    for cur.Next.Next != nil {
        cur = cur.Next
    }
    node := cur.Next
    cur.Next = nil
    node.Next = head
    
    return node
}
上面的提交超时了，时间复杂度是 O(n*k)
怎么变成线性的时间复杂度呢？

2. 看题解：
将链表先连成环，然后计算出链表的长度
然后在新的有环链表中，找到新链表中的最后一个结点第（n-1）- (k mod n)个结点，然后断开
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil {
        return head
    }
    n := 1
    iter := head
    for iter.Next != nil {
        iter = iter.Next
        n++
    }
    add := n - k%n
    if add == n {
        return head
    }
    iter.Next = head
    for add > 0 {
        iter = iter.Next
        add--
    }
    ret := iter.Next
    iter.Next = nil
    return ret
}
细节点在于找断开的地方，比较容易出错

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 如果想要减少时间复杂度一般都需要数学知识去运用是使用的
