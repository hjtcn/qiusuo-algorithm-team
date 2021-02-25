/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。
示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

1. Clearfication:

想要反转 m 到 n 的链表，那么怎么反转整个链表呢？
如：
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
怎么反转单链表呢？
我想到了反转数组的时候，进行下标互换 
ret[i],ret[j] = ret[j],ret[i]
链表里面是怎么 reverse 的呢？
发现自己没有思路哈
2. Coding:

3. 看题解:

方法一：迭代：反转整个链表
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    
    return prev
}

方法二：递归: 整个链表的反转：
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    last := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    
    return last
}

递归算法：明确递归函数的定义
我们的 reverse 函数定义是这样的：
输入一个节点head, 将以 head 为起点的链表反转，并返回反转以后的头结点
不要跳进递归（你的脑袋能压几个栈呀？）
反转链表的前N个节点：

var successor *ListNode
func revserseN(head *ListNode,n int)*ListNode {
    if n == 1 {
        successor = head.Next
        return head
    }
    last := revserseN(head.Next,n - 1)
    head.Next.Next = head
    head.Next = successor
    return last
}

1. base case 变为 n == 1,反转一个元素，就是它本身，同时要记录后驱节点
2. 上面的递归代码直接把 head.Next 设置为了 nil,因为整个链表反转后原来的 head 变成了整个链表的最后一个节点。现在head节点在递归反转之后不一定是最后一个节点了，所以要记录后驱 successor（第n+1个节点），反转之后将 head 连接上。

反转链表的一部分：
如果 left == 1,相当于反转链表开头的right 个元素，也就是上面的反转前 right 个元素

 var successor *ListNode
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == 1 {
        return revserseN(head, right)
    }
    head.Next = reverseBetween(head.Next, left - 1,right - 1)
    return head
}

func revserseN(head *ListNode,n int)*ListNode {
    if n == 1 {
        successor = head.Next
        return head
    }
    last := revserseN(head.Next,n - 1)
    head.Next.Next = head
    head.Next = successor
    return last
}

还是有一点不理解的地方
head.Next = reverseBetween(head.Next, left - 1,right - 1)
执行 left - 1 和 right - 1 这一部分没有想清楚是为什么
反转链表II 迭代法：
核心点在于 记录反转前的节点，然后将反转后的节点和 反转前的一个节点和最后一个节点连接起来

funv reverseBetween(head *ListNode, m int,n int) *ListNode {
        prevM := &ListNode{Next: head}
    
    for i := 1;i < m;i++ {
            prevM = prevM.Next
    }
    
    first := prevM.Next
    second := first.Next
    if  second == nil {
         return head
    }
    
    three := second.Next
    for i := 0;i < n - m;i++ {
         second.Next = first
       first,second = second,three
       if three != nil {
            three = three.Next
       }
    }
    
    prvM.Next.Next,prevM.Next = second,first
    
    if m == 1 {
         return prevM.Next
    }
    
    return head
}

4. 复杂度分析:
时间复杂度：O(n)
空间复杂度：递归:O(n) 迭代：O(1)

5. 总结:
5.1： 不要跳进递归，脑子里面存不了几个栈
5.2:   思维是一点点的完善的，不要一下子跳跃太快
