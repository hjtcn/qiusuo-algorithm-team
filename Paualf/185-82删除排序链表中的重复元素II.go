存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
返回同样按升序排列的结果链表。
示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]
 
提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

    1. Clarification:
        一开始想到使用快慢指针，因为在链表中要删除元素，需要记录前一个节点的位置
            1 -> 2 -> 3 -> 3
            slow
            high
            high ++ ,然后判断是否与slow相等，不想等的时候，slow.Next = high, 但是这个时候走到3的时候不对了，slow=high=3,3已经加进去了，所以这个思路不对
            需要再加一个指针
            dummy := ListNode {
            }
            dummy.Next = head
            pre := dummy
            slow := pre.Next
            high := slow.Next
            pre 指向第0个节点 -> 头节点 为了删除这种第一个节点就是重复元素的位置 1->1->1
            slow 指向第1个节点 ->
            high 执行第2个节点 -> 
func deleteDuplicates(head *ListNode) *ListNode {
    // 这个判断完以后最少🈶️2个节点
    for head == nil {
        return head
    }
    // 如果一开始就有重复元素呢？
    
     dummy := &ListNode {
     }
     pre := dummy
     dummy.Next = head
     slow := pre.Next
     for slow != nil {
         fmt.Println(slow.Val)
         high := slow.Next
         flag := false
         // 有重复元素
         for slow != nil && high != nil && slow.Val == high.Val {
            flag = true
            high = high.Next
         }
         if flag {
             // 细节点，这个时候high 已经不相等了
            slow = high
         }else {
             // 当有重复节点的时候才进行移动
             // 更改 pre,移动 slow 和 pre
            pre.Next = slow
            slow = slow.Next 
            pre = pre.Next
         }
     }
    // 不能返回head，head这个时候没有变
     return dummy.Next
}
一开始就有重复元素的时候上面的代码运行就不符合预期了

2. 看题解：
cur 指向哑节点，随后开始对链表进行遍历。如果当前 cur.Next 与 cur.Next.Next 对应的元素相同，那么我们就需要将 cur.Next 以及后面拥有相同元素值的链表节点全部删除。
记下这个元素值为x,随后不断将 cur.Next 从链表中移除，直到 cur.Next 为空节点或者元素值不等于x为止。这个时候，我们将链表中所有元素等于x的节点全部删除。
如果当前 cur.Next 与 cur.Next.Next 不相同，说明链表中只有一个元素为 cur.Next  的节点，就可以将 cur 指向 cur.Next
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{0, head}
    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
            x := cur.Next.Val
            for cur.Next != nil && cur.Next.Val == x {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }
    return dummy.Next
}

使用hash表记录也是可以的
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    //傻瓜解法查找表 + 虚拟节点
    hash := make(map[int]int)
    dummy := &ListNode{Next:head}
    for head != nil{ //先建立查找表
        hash[head.Val]++
        head = head.Next
    } 
    pre := dummy //依靠pre来删除节点
    head = dummy.Next
    for head != nil{
        if hash[head.Val] > 1{//满足条件，直接删除该节点
            pre.Next = head.Next
        }else{
            pre = pre.Next //若不是目标节点，则pre也需要前进
        }
        head = head.Next
    }
    
    return dummy.Next //返回
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 感觉官方题解的思路还是蛮好的，但是比较难想到，把重复的移除，每个元素都检查一遍，然后边界条件处理的也很nice
