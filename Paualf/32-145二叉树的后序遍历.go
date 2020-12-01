/*
给定一个二叉树，返回它的 后序 遍历。
示例:
输入: [1,null,2,3]  
   1
    \
     2
    /
   3 
输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

Clearfication:
二叉树的后序遍历:
左节点-> 根节点 -> 右节点

预估复杂度：
时间复杂度：O(n)
空间复杂度：O(height) 二叉树的高度，也就是递归栈的深度

Solution:
1. 递归进行遍历：
2. 使用辅助栈空间进行迭代，迭代的时候主要

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root, &ret)
    return ret
}
func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    helper(root.Right,ret)
    *ret = append(*ret, root.Val)
}

helper的参数写的还是不熟练：
ret *[]int

迭代的时候，还是需要想清楚什么时候将节点放入到栈中，然后什么时候弹出栈，然后进行遍历, 自己还是没有想清楚，看了题解
放入数组的关键条件:
想要弹出这个节点，左节点和右节点都要访问过，左节点是一直放进去的节点，就不需要判断了，那么我们的关注点就放在右节点上了
root.Right == nil 
右节点为空或者已经访问过，这个时候该节点就可以取出来了

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode)[]int {
    stack := []*TreeNode{}
    ret := []int{}
    var lastVisit *TreeNode
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 看一下栈顶元素不弹出来
        root := stack[len(stack)-1]
        if root.Right == nil || root.Right == lastVisit {
            // pop
            stack := stack[:len(stack)-1]
            ret = append(ret, node.Val)
            root = nil
        }else {
            root = root.Right
        }
    }
    return ret
}

题解是弹出来，判断一下，如果不需要，再放进去

func postorderTraversal(root *TreeNode) (res []int) {
    stk := []*TreeNode{}
    var prev *TreeNode
    for root != nil || len(stk) > 0 {
        for root != nil {
            stk = append(stk, root)
            root = root.Left
        }
        root = stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        if root.Right == nil || root.Right == prev {
            res = append(res, root.Val)
            prev = root
            root = nil
        } else {
            stk = append(stk, root)
            root = root.Right
        }
    }
    return
}

可以先看一下栈顶元素，不弹出来，然后如果需要弹出来，然后再弹出来

func postorderTraversal(root *TreeNode) (res []int) {
    stk := []*TreeNode{}
    var prev *TreeNode
    for root != nil || len(stk) > 0 {
        for root != nil {
            stk = append(stk, root)
            root = root.Left
        }
        root = stk[len(stk)-1]
        if root.Right == nil || root.Right == prev {
            stk = stk[:len(stk)-1]
            res = append(res, root.Val)
            prev = root
            root = nil
        } else {
            root = root.Right
        }
    }
    return
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n): 递归使用栈空间大小 或者 迭代开辟数组空间大小

总结：
1. 迭代法和前序遍历和中序遍历有不一样的地方，主要还是节点弹出的顺序不一样

2. 我们可以将三道题的迭代放到一起，对比分析，举一反三
