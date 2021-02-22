/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。
示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
1. Clearfication:
看书印象中刚看过这道题，现在去写这道题，还是没有思路，不知道如何下手去做，就去看题解去了

2. Coding
     没有啥思路

3. 看题解
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    
    start := &ListNode{
        0,
        nil
    }
    write,cur := start,start
    
    for i := 0;i < m;i++ {
        tmp := head.Next
        cur.Next,head.Next = head,cur.Next
        write,cur,head = cur,cur.Next,tmp
    }
    
    for i := 0;i < n - m;i++ {
        tmp := head.Next
        write.Next,head.Next = head,write.Next
        head = tmp
    }
    
    if head != nil {
        cur.Next = head
    }
    
    return start.Next
}

1. 先把问题简化成前N个数进行反转，得到reverseListN
2. 再判断 m == 1的情况，直接使用reverseListN
3. m != 1时，指针后移，m和n同时-1

var temp *ListNode
func reverseListN(head *ListNode,N int)*ListNode {
    if N == 1 {
        temp = head.Next
        return head
    }
    p := reverseListN(head.Next, N - 1)
    head.Next.Next = head
    head.Next = temp
    return p
}
func reverseBetween(head *ListNode,m,n int)*ListNode {
    if m == 1 {
        return reverseListN(head,n)
    }
    head.Next = reverseBetween(head.Next,m-1,n-1)
    return head
}


func reverseBetween(head *ListNode,m,n int) *ListNode {
    dummy := &ListNode{}
    head,dummy.Next = dummy,head
    
    for i := 0;i < m - 1;i++ {
        dummy = dummy.Next
    }
    
    node := dummy.Next
    for i := 0;i < n - m;i++ {
        dummy.Next,node.Next,node.Next.Next = node.Next,node.Next.Next,dummy.Next
    }
    
    return head.Next
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

5. 总结：
5.1：对链表操作不是很熟悉，需要多多联系
