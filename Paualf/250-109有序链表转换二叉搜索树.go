给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

1. Clarification:
高度平衡而且原链表是有序的

将链表元素放入到数组中，题目就变成了将有序数组转化为二叉搜索树

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    arr := []int{}

    for head != nil {
        arr = append(arr,head.Val)
        head = head.Next
    }

    return toBST(arr)
}

func toBST(arr []int) *TreeNode {
    length := len(arr)
    if length <= 0 {
        return nil
    }

    mid := length / 2
    node := &TreeNode{
        Val: arr[mid],
    }
    node.Left = toBST(arr[:mid])
    node.Right = toBST(arr[mid+1:])

    return node
}

2. 看题解：
分治使用快慢指针找到链表中点，然后进行处理

中序遍历定义一个全局指针，然后移动这个指针

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 思路和想法还是蛮重要的
