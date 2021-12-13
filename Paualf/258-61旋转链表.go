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

1. Clarification:
旋转链表，找到最后一个节点，将最后一个节点的位置指向头节点
问题：转换，怎么找到最后节点

curr := head
var last *ListNode
for curr.Next.Next != nil {

}
当 curr.Next.Next == nil 的时候退出，这个时候，curr.Next 是最后一个节点
last := curr.Next
last.Next = head
curr.Next = nil

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    for k > 0 {
       head = move(head)
        k--
    }

    return head
}

// 移动一次,返回移动一次以后的头节点信息
func move(head *ListNode) *ListNode {
    curr := head
    var last *ListNode
    last = head

    if curr.Next != nil {
        last = curr.Next
    }
    
    for curr.Next != nil &&  curr.Next.Next != nil {
        last = curr.Next.Next
        curr = curr.Next
    }

    if last != head {
        // 断开
        curr.Next = nil
        // 尾指针指向头指针
        last.Next = head
        //fmt.Println(last.Val)
    }

    return last
}
有思路的情况下，很多边界的case还是没有想清楚。。。】

如果执行
[1,2,3] 2000000000
会超时，可以结算下长度，然后用次数 % 长度 = 周期

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    length := getListLength(head)

    cycle := k % length

    for cycle > 0 {
       head = move(head)
        cycle--
    }

    return head
}

func getListLength(head *ListNode) (length int) {
    for head != nil {
        length += 1
        head = head.Next
    }
    return
}

// 移动一次,返回移动一次以后的头节点信息
func move(head *ListNode) *ListNode {
    curr := head
    var last *ListNode
    last = head

    if curr.Next != nil {
        last = curr.Next
    }

    for curr.Next != nil &&  curr.Next.Next != nil {
        last = curr.Next.Next
        curr = curr.Next
    }

    if last != head {
        // 断开
        curr.Next = nil
        // 尾指针指向头指针
        last.Next = head
        //fmt.Println(last.Val)
    }

    return last
}
2. 看题解：
闭合为环

链表为环，然后，新链表的最后一个节点为原链表的第 (n-1） - (k mod n) 个节点（从0开始计数）

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1：技术方案的细节点以及是否想的清楚是这个东西能做出来的关键
