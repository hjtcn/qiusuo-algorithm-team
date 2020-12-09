/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
 
示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/

Clearfication:
找重复性：root.Left,root.Right = root.Right,root.Left
重复该步骤，然后继续遍历 
mirrorTree(root.Left)
mirrorTree(root.Right)

  复杂度预估：
时间复杂度：O(n)： 遍历每个节点
空间复杂度:  O(n) 递归栈用到的空间

Coding 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return
    }
    root.Left,root.Right = root.Right,root.Left
    mirrorTree(root.Left)
    mirrorTree(root.Right)   
}

出现报错了：
Line 16:Char 1: missing return at end of function (solution.go)
看到了函数还有返回值，这就有点懵了
那么我们加一个 helper 帮我们处理呢

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    helper(root)
    return root
}

func helper(root *TreeNode) {
    if root == nil {
        return
    }
    root.Left,root.Right = root.Right,root.Left
    helper(root.Left)
    helper(root.Right)
}

然后提交就通过了
使用 queue 层次遍历交换左右节点
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    helper(root)
    return root
}

func helper(root *TreeNode) {
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        node.Left,node.Right = node.Right,node.Left
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}

看了题解以后，自己写的递归是自顶向下的递归，那么自底向上的递归是什么呢
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    helper(root)
    return root
}

func helper(root *TreeNode) {
    if root == nil {
        return
    }
    
    helper(root.Left)
    helper(root.Right)
    root.Left,root.Right = root.Right,root.Left
}

最后直接可以return root 的，自己还加了一个函数。。。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := mirrorTree(root.Left)
    right := mirrorTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

总结：
1. 这道题第二次做，是有不一样的感受的

2. 找重复性，树的问题就是解决重复子问题最多的数据结构
