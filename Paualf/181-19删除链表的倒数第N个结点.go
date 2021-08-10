给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
进阶：你能尝试使用一趟扫描实现吗？
示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：
输入：head = [1], n = 1
输出：[]
示例 3：
输入：head = [1,2], n = 1
输出：[1]
 
提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

 1. Clarification:
        一次扫描的方法没有想出来
        那就先扫描2次吧：
            求出 length 
            node := head
            for i := 0;i < length - n;i++ {
                node = node.Next
            }
            node.Next = node.Next.Next
            return head

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    node := head
    for node != nil {
        length++
        node = node.Next
    }
    fmt.Println(length)
    node = head
    
    for i := 0;i < length - n - 1;i++ {
        node = node.Next
    }
    
    node.Next = node.Next.Next
    return head
}

上面设计的有问题，应该是 < length - n - 1
遇到只有一个元素的时候 [1] 1, 的时候报错了，这个算法不适用于头节点的删除
加入虚拟头节点进行解除：

/*
    1. Clarification:
        一次扫描的方法没有想出来
        那就先扫描2次吧：
            求出 length 
            node := head
            for i := 0;i < length - n;i++ {
                node = node.Next
            }
            node.Next = node.Next.Next
            return head
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    newNode := &ListNode{
        Val:0,
        Next:head,
    }
    length := 0
    node := newNode
    for node != nil {
        length++
        node = node.Next
    }
   // fmt.Println(length)
    node = newNode
    for i := 0;i < length - n - 1;i++ {
        node = node.Next
    }
    node.Next = node.Next.Next
    return newNode.Next
}

加了头节点以后，for循环的遍历的时候 是不是要改成 < length - n 了，试了下，不可以哈，因为是倒着数的，它的相对位置是没有变的
还有最后返回的时候要返回新建链表的下一个结点位置

2. 看题解：
怎么一次循环就可以搞定的哈？
双指针的经典应用：
如果要删除倒数第n个节点，让fast移动n步，然后让fast和slow同时移动，直到fast指向链表末尾。删掉slow所指向的节点就可以了
官方题解写的还是蛮精简的，里面有细节在里面，first指向的是head,移动了n步，如果指向dummpy，需要移动n+1步

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    first, second := head, dummy
    for i := 0; i < n; i++ {
        first = first.Next
    }
    for ; first != nil; first = first.Next {
        second = second.Next
    }
    second.Next = second.Next.Next
    return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    first, second := dummy, dummy
    for i := 0; i < n + 1; i++ {
        first = first.Next
    }
    for ; first != nil; first = first.Next {
        second = second.Next
    }
    second.Next = second.Next.Next
    return dummy.Next
}

3. 复杂度分析：
时间复杂度：O(n) : 遍历链表计算长度
空间复杂度：O(1)

4. 总结：
4.1: dummy 节点为什么要加，方便删除头节点，删除操作统一

4.2: 双指针的应用，数轴上的数学题，好理解，不好想到。。。
