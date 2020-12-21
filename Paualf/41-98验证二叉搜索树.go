/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

Clearfication:
二叉搜索树的特点：
左子树小于当前节点
右子树大于当前节点
所有左子树和右子树也满足这个特点

局部 -> 整体
以小见大 -> 推而广之
想到的解法：

中序遍历二叉搜索树，返回一维数组
然后对一维数组进行判断：
输入：一维数组
返回： 如果一维数组是递增的，则返回true，否则返回false

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    ret := []int{}
    inorder(root, &ret)
    return isIncrement(ret)
}
func inorder(root *TreeNode,ret *[]int) {
    // terminator
    if root == nil {
        return
    }
    inorder(root.Left,ret)
    *ret = append(*ret, root.Val)
    inorder(root.Right,ret)
}
func isIncrement(arr []int)bool {
    for i := 0;i < len(arr) - 1;i++ {
        if arr[i] >= arr[i + 1] {
            return false
        }
    }
    return true
}

第一遍提交的时候报错了
isIncrement 里面少加了 = 的判断

> -> >=

func isIncrement(arr []int)bool {
    for i := 0;i < len(arr) - 1;i++ {
        if arr[i] > arr[i + 1] {
            return false
        }
    }
    return true
}

二叉树的中序遍历: 
使用栈进行迭代：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    ret := []int{}
    inorder(root, &ret)
    return isIncrement(ret)
}
func inorder(root *TreeNode,ret *[]int) {
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        *ret = append(*ret, node.Val)
        
        root = node.Right
    }
}
func isIncrement(arr []int)bool {
    for i := 0;i < len(arr) - 1;i++ {
        if arr[i] >= arr[i + 1] {
            return false
        }
    }
    return true
}

看到题解的时候中间不需要使用数组来存储元素去判断数组是否为递增的
在中序遍历的中间就可以判断是否可以返回 true/false

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    
    return true
}

对比自己写的，它使用了root节点，自己的代码中引入了node节点，少使用了一个遍历，而且看着特别舒服
递归的思路是自己没有想到的，看到题解才看到了

func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode,lower,upper int) bool {
    if root == nil {
        return true
    }
    
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    
    return helper(root.Left, lower, root.Val) && helper(root.Right,root.Val,upper)
}

感觉思路还是利用二叉搜索树的性质，不断变更左右边界，去比较，看是否满足
复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度: O(n) 递归调用栈开辟空间 or 使用队列存储树指针元素开辟空间

总结：
1. 要了解理论知识，二叉搜索树的性质，你才能理解和写出来代码

2. 利用性质不断变换左右边界进行递归
