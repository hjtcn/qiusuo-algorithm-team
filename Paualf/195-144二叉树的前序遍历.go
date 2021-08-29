给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

 

示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[1,2]
示例 5：


输入：root = [1,null,2]
输出：[1,2]
 

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
 

进阶：递归算法很简单，你可以通过迭代算法完成吗？

1. Clarification:
前序遍历：
前中后

递归的话还是分解子问题:
Terminator

Process current Logic

Drill down left and right
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    *ret = append(*ret, root.Val)
    helper(root.Left,ret)
    helper(root.Right,ret)
}

如果用迭代怎么去解决呢？

想了下，没有什么思路

2.看题解：
递归的时候隐式地维护了一个栈：需要显式地将这个栈模拟出来
func preorderTraversal(root *TreeNode) []int {
    vals := []int{}
    
    stack := []*TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            vals = append(vals, node.Val)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    return vals
}
巧妙的把这个处理过程模拟出来了，还是蛮厉害的

根左右


3.复杂度分析：
时间复杂度：O(n) : n个节点
空间复杂度：O(n) : 调用栈空间 or 开辟栈所用空间

总结：
4.1: 迭代的时候没有想出来，思路还是蛮巧妙的，特别是 for循环的时候，里面的变换，还是蛮细节的
