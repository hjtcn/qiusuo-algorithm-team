/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

1. Clearfication:
数据结构本质上就是定义这种性质，然后维护这种性质
二叉树中任意节点的左右子树的深度相差不超过1
最直接的就是计算每个节点左子树的深度和右子树的深度，判断两个深度是否超过1
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    leftHeight = height(root.Left)
    rightHeight = height(root.Right)
    
    if leftHeight - rightHeight > 1 || rightHeight - leftHeight > 1 && isBalanced(root.Left) && isBalanced(root.Right) {
        return fasle
    }
    
    return true
}
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHight = height(root.Left) + 1
    rightHeight = height(root.Right) + 1
    
    return max(leftHight,rightHeight) + 1
}
func max(a,b int) {
    if a > b {
        return a
    }
    
    return b
}

上面代码是有问题的，首先计算height的地方没有想清楚，应该是计算左节点的高度，然后计算右节点的高度，然后加上本层节点，举个例子，如果是子节点，左节点高度为0，右节点高度为0，那么这个节点的高度为1，所以上面的代码明显是有问题的
还有我们既然要使用root.Left，root.Right 那么我们要判断 root.Left 和 root.Right 是否存在，同时我们还要判断root是否存在，自己还是将 root == nil 的判断给漏掉了 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    leftHeight := height(root.Left)
   
    rightHeight := height(root.Right)    
   
    if leftHeight - rightHeight > 1 || rightHeight - leftHeight > 1 || !isBalanced(root.Left) || !isBalanced(root.Right) {
        return false
    }
    
    return true
}
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHight := height(root.Left) 
    rightHeight := height(root.Right)
    
    return max(leftHight,rightHeight) + 1
}
func max(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}

上面的解法是自顶向下的解法，时间复杂度很高，2^n 的时间复杂度，因为有重复计算，我们用递归也是这个道理，里面有重复计算，时间复杂度会非常高
那么自底向上的解法我们怎么写呢？
自底向上，我们怎么将节点移动到最下面去判断深度是否相差1呢
没想出来，看了题解：

func isBalanced(root *TreeNode) bool {
    return height(root) >= 0
}
func height(root *TreeNode)int {
    if root == nil {
        return 0
    }
    leftHeight := height(root.Left)
    rightHeight := height(root.Right)
    if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
        return -1
    }
    
    return max(leftHeight,rightHeight) + 1
}
func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}
func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}

看完题解以后还是感觉蛮巧妙的，-1 这个点用的比较好感觉
然后时间复杂度也降下来了，所有的地方只是计算一次，时间复杂度变成了O(n)
第一个的时间复杂度自己分析错了，想成了递归，其实是每一个节点都要计算它的高度，height(root), 计算高度的时间复杂度是O(log(N)),所以时间复杂度平均来说是nlogn(n),如果二叉树最坏情况下是一个链表的话，它的计算高度最坏是O(n), 这个时间的时间复杂度是O(n*n)

复杂度分析：
自顶向下：时间复杂度：O(n*n） 空间复杂度：O(n) 
自底向上：时间复杂度：O(n) 空间复杂度：O(n)
自顶向下的时间复杂度高是在于重复计算height(root)

总结：
1. 二叉树计算高度的代码还是不熟悉

2. 自底向上的思维和想法还是不熟悉
