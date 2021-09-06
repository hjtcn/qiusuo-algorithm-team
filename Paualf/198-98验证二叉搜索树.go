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

1. Clarification:

验证二叉搜索树

二叉搜索树的特点是：中序遍历为升序

那怎么在中序遍历的时候判断呢？

func isValidBST(root *TreeNode) bool {

}

func helper(root *TreeNode,lastVal int) {
    if root == nil {
        return
    }
    helper(root.Left)
    if root.Val < lastVal {
        return false
    }
    helper(root.Right)
}
没有想清楚
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
    helper(root,&ret)

    for i := 0;i < len(ret) - 1;i++ {
        if ret[i] >= ret[i+1] {
            return false
        }
    }
    return true
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret,root.Val)
    helper(root.Right,ret)
}
>= 这里自己写成了等于

2. 看题解：

看了题解的第一种做法：收缩界限
第2种做法：使用了迭代的方法去判断，我一开始想用递归去判断，没有想清楚

func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
    if root == nil {
        return true
    }
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}


3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 知道不一定可以做到，如果做不到是不是还是不知道哈

4.2: 写代码这种东西，光有思路还是不行，你要动动手，将思路写出来才行，如果仅仅是想法，其实没啥用
