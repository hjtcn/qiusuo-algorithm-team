给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
 
示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]
 
提示：
链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105

1. Clarification:

Q:  如果是自己去做链表排序会怎么做呢？
A:  

1.1: 将链表遍历放到数组中
1.2: 对数组进行快速排序
1.3: 将排序后的数组放入到新建的链表当中

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    arr := []int{}
    
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    sort.Ints(arr)
    
    //fmt.Println(arr)
    
    head = &ListNode{Val:arr[0]}
    p := head
    for i := 1;i < len(arr);i++ {
        tmp := &ListNode{
            Val:arr[i],
        }
        p.Next = tmp
        p = p.Next
    }

    return head
}

2. 看题解:
方法一：自顶向下归并排序：
找到链表的中点，以中点为界，将链表拆分成两个子链表
对两个子链表分别排序
将两个排序后的子链表合并，得到完整的排序后的链表

感觉比较巧妙的是 sort 函数的tail参数还是蛮特别的
func sort(head, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
            fast = fast.Next
        }
    }

    mid := slow
    return merge(sort(head, mid), sort(mid, tail))
}
自底向上归并排序

用subLength 表示每次需要排序的子链表的长度，初始时 subLength =1 
每次将链表拆分成若干个长度为subLength 的子链表，按照每两个子链表一组进行合并，合并后即可得到若干个长度为subLength x 2的有序子链表
将subLength 的值加倍，重复第2步，对更长的有序子链表进行合并操作，直到有序子链表的长度 大于或等于length,整个链表排序完毕

自底向上
func merge(head1, head2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    length := 0
    for node := head; node != nil; node = node.Next {
        length++
    }

    dummyHead := &ListNode{Next: head}
    for subLength := 1; subLength < length; subLength <<= 1 {
        prev, cur := dummyHead, dummyHead.Next
        for cur != nil {
            head1 := cur
            for i := 1; i < subLength && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }

            var next *ListNode
            if cur != nil {
                next = cur.Next
                cur.Next = nil
            }

            prev.Next = merge(head1, head2)

            for prev.Next != nil {
                prev = prev.Next
            }
            cur = next
        }
    }
    return dummyHead.Next
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 好的方法和思路还是要花时间的进行思考和实验的，不要伪勤奋

