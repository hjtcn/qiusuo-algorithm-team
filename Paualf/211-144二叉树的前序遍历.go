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
根左右
递归算法很简单，你可以想出来迭代算法么？
思路: 迭代算法使用stack 去模拟递归的过程

递归：
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

2. 看题解：
分析了下题解的过程：
先向左，如果有元素加入到结果数组及栈中
从栈中弹出元素看它右边节点是否存在

如果是我实现的话，比较难考虑的是第2步
还有这个过程怎么反复执行呢？

func preorderTraversal(root *TreeNode)[]int {
    ret := []int{}
    
    stack := []*TreeNode{}
    node := root
    
    for node !=nil || len(stack) > 0 {
        for node != nil {
            ret = append(ret,node.Val)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack) - 1].Right
        stack = stack[:len(stack)-1]
    }
    
    return ret
}
比较难想和不好解决的地方在于 for node !=nil || len(stack) > 0的这个判断条件

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4.总结：
4.1: 以前递归都写不出来，现在递归可以写出来了，难的还是迭代哈
