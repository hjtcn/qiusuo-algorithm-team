给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。

删除完毕后，请你返回最终结果链表的头节点。

 你可以返回任何满足题目要求的答案。
（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：
输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。

示例 2：
输入：head = [1,2,3,-3,4]
输出：[1,2,4]

示例 3：
输入：head = [1,2,3,-3,-2]
输出：[1]
 
提示：
给你的链表中可能有 1 到 1000 个节点。
对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.

1. Clarification:

Q: 想要删除元素
A: 需要有一个指针记录 pre

Q: 要向后遍历
A: 要有一个指针向后移动 curr

Q: 要向后遍历，并且退出
A: 要有一个for循环遍历 for curr != nil

Q: 要统计和
A: 要有一个 currentSum 记录 

Q: currtSum 变化
A: 是0 的时候加上以后，下一次遍历就不为0了

Q: 删除元素
A: pre.Next = curr.Next

Q: 关于中间出现的加到一起等于0的情况，这部分怎么处理呢？
A: 这部分没有想好。。。


2. 看题解：
前缀和：
我们可以考虑如果给的入参不是链表是数组的话，只需要求出前缀和，对于前缀和相同的项，那他们中间的部分即是可以消除掉的，比如以 [1, 2, 3, -3, 4] 为例，其前缀和数组为 [1, 3, 6, 3, 7] ，我们发现有两项均为 3，则 6 和 第二个 3 所对应的原数组中的数字是可以消掉的。换成链表其实也是一样的思路，把第一个 3 的 next 指向第二个 3 的 next 即可
作者：philhsu
链接：https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solution/c-jian-ji-dai-si-lu-by-philhsu/

func removeZeroSumSublists(head *ListNode) *ListNode {
    dummy := &ListNode{Val : 0, Next : head}
    sum := 0
    dict := make(map[int]*ListNode)
    for p := dummy; p != nil; p = p.Next {
        sum += p.Val
        dict[sum] = p
    }
    sum = 0
    for p := dummy; p != nil; p = p.Next {
        sum += p.Val
        p.Next = dict[sum].Next
    }
    return dummy.Next
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 前缀和，好戏那个第一次有意识的看到这个，没有往这边想过，感觉还是数学哈这个
