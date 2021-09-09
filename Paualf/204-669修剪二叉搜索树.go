给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

 

示例 1：


输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
示例 2：


输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]
示例 3：

输入：root = [1], low = 1, high = 2
输出：[1]
示例 4：

输入：root = [1,null,2], low = 1, high = 3
输出：[1,null,2]
示例 5：

输入：root = [1,null,2], low = 2, high = 4
输出：[2]
 

提示：

树中节点数在范围 [1, 104] 内
0 <= Node.val <= 104
树中每个节点的值都是唯一的
题目数据保证输入是一棵有效的二叉搜索树
0 <= low <= high <= 104

1.Clarification:
Q: 怎么分解子问题？
A: 如果这个节点 < low || 节点的值 > high 这个节点就要删除

Q: 怎么删除呢？
这个逻辑没有想清楚。。。

2. 看题解：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
    	return
    }
    if root.Val > high {
        return trimBST(root.Left,low,high)
    }
    if root.Val < low {
        return trimBST(root.Right,low,high)
    }
    
    root.Left = trimBST(root.Left,low,high)
    root.Right = trimBST(root.Right,low,high)
    
    return root
}
这个过程还是有点绕的

如果 root.Val > high 去左边找节点，因为右边的节点都不符合要求
如果 root.Val < low 的时候去右边找节点，因为左边的节点都比较小，不符合要求

3.复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归调用栈空间

总结：
4.1: 做了这么多二叉树的题感觉还是不是很熟悉，是不是自己理论知识少了
4.2: 也有可能是理论知识知道，但是还是不会，因为还不是很熟悉
