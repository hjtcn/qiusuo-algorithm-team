给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
示例 1：
输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。

示例 2：
输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true

示例 3：
输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。
 
提示：
二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
链表包含的节点数目在 1 到 100 之间。
二叉树包含的节点数目在 1 到 2500 之间。

1. Clarification:
思路：
    1.1: 遍历二叉树中的节点
    1.2: 判断是否 == 链表中的值，curr := head
    1.3: 如果相等，则继续向下走，如果 curr == nil,则退出
 
 2. 看题解：
枚举二叉树中的每个节点为起点往下的路径是否有与链表相匹配的路径。设计一个递归函数 dfs(rt,head): rt表示当前匹配到的二叉树节点，head 表示当前匹配到的链表节点，整个函数返回布尔值表示是否有一条该节点往下的路径与head节点开始的链表匹配，如匹配返回true,否则返回false，分四种情况：

链表已经全部匹配完，匹配成功,返回true
二叉树访问到了空节点，匹配失败，返回false
当前匹配的二叉树上节点的值与链表节点的值不想等，匹配失败，返回false
前3种情况都不满足，说明匹配成功了一部分，我们需要继续递归匹配 dfs(rt.Left,head.Next)  dfs(rt.Right,head.Next)

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
func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }
    return dfs(root,head) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func dfs(root *TreeNode,head *ListNode) bool {
    if head == nil {
        return true
    }

    if root == nil {
        return false
    }
    
    if root.Val != head.Val {
        return false
    }

    return dfs(root.Left,head.Next) || dfs(root.Right,head.Next)
}

3. 复杂度分析：
时间复杂度：O（n*n) 
空间复杂度：O(n)

4. 总结：
4.1: 关注点还是在怎么把一个复杂的问题分解为一个小问题去解决，fail fast,然后去处理正常的情况
