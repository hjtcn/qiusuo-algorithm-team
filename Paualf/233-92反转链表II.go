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
 
1. Clarification:
Q: 和昨天做的反转链表有什么区别呢？
A: 个人感觉最大的区别在于反转后怎么链接到前面的节点和后面的节点

Q: 什么是前面的节点和后面的节点呢？
A: 比如说 1 2 3 4 5  如果我 反转链表 找到 pre = 1, next = 5, 然后 在 2 3 4 的时候断开，去反转  2 3 4

Q: 反转完成 2 3 4 以后，怎么拼回去呢？
A: 1 4 3 2 5

代码逻辑：
找到分开的节点，比如题目中的 1 5
将中间的节点进行反转
将节点连接到一起

内存溢出了
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var pre,next *ListNode
    curr := head
    i := 1

    for curr != nil {
        if i + 1 == left {
            pre = curr
        }

        if i == right + 1 {
            next = curr
            // 断开
            curr.Next = nil
        }

        curr = curr.Next
        i++
    }

   tmp := pre.Next
   pre.Next = nil

   pre.Next = reverseList(pre.Next)
   tmp.Next = next

   return head
}

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }

    return head
}
感觉断开然后连接的代码逻辑不清晰，所以会出错的

2. 看题解：
看完题解以后发现自己写的代码真的有点水，逻辑不够清晰
func reverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = pre
        pre = curr
        curr = next
    }

    return head
}

func reverseBetween(head *ListNode, left, right int)*ListNode {
    dummyNode := &ListNode{Val: -1}
   	dummyNode.Next = head
    
    pre := dummyNode
    // 第1步：从虚拟节点走 left - 1步，走到left 节点的前一个节点
    for i := 0;i < left - 1;i++ {
    	pre = pre.Next
    }
    
    // 第2步，从pre再走 right - left + 1，来到 right 节点
    rightNode := pre
    for i := 0;i < right - left + 1;i++ {
    	rightNode = rightNode.Next
    }
    
    // 第 3 步：切断出一个子链表(截取链表)
    leftNode := pre.Next
    curr := rightNode.Next
    
    // 注意：切断连接
    pre.Next = nil
    rightNode.Next = nil
    
    // 第4步，反转链表的子区间
    reverseLinkedList(leftNode)
    
    pre.Next = rightNode
    leftNode.Next = curr
    return dummyNode.Next
}
穿针引线
3个变量的不断倒腾
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
4.1: 逻辑清晰要写注释，因为代码是要维护下去的哈
