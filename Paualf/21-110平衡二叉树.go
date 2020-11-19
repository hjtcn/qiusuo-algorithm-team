/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：true
示例 2：
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
示例 3：
输入：root = []
输出：true
 
提示：
树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104
平衡二叉树：
一个二叉树的每个节点的左右两个子树的高度差的绝对值不超过1
*/

Wrong code : 没有考虑子问题
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
     if leftHeight > rightHeight + 1 || rightHeight > leftHeight + 1 {
         return false
     }
     return true
}
func height(root *TreeNode)int {
    if root == nil {
        return 0
    }
    left := height(root.Left)
    right := height(root.Right)
    return max(left, right) + 1
}
func max(a,b int)int {
    if a > b {
        return a
    }
    return b
}

怎么组织代码又卡这里了。。。
卡在了没有处理重复子问题上面，看了题解加了两个对于子问题的判断

isBalanced(root.Left) 和 isBalanced(root.Right) 的判断

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
     if leftHeight > rightHeight + 1 || rightHeight > leftHeight + 1 || !isBalanced(root.Left) || !isBalanced(root.Right) {
         return false
     }
     return true
}
func height(root *TreeNode)int {
    if root == nil {
        return 0
    }
    left := height(root.Left)
    right := height(root.Right)
    return max(left, right) + 1
}
func max(a,b int)int {
    if a > b {
        return a
    }
    return b
}

复杂度分析：
时间复杂度：大于O（n) 的，因为中间有重复计算的逻辑。看了题解写了O(N*N),一开始没有理解，后来看到别人这样分析，还是理解了一点
对于节点node,高度函数的时间复杂度与其高度有关，平均情况下高度h=logN,最坏情况下，树为斜树，高度函数的复杂度为O(n)
所以总的时间复杂度是O(N^2)
空间复杂度： O(n):树的高度
 
自底向上递归：
这个是参考题解的，自己卡在的点在于如何判断返回值 也就是 return height(root) >= 0
height 函数不仅仅是计算高度，而且还是负责主函数的返回值，这个地方还是需要思考的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
func max(x,y int) int{
    if x > y {
        return x
    }
    return y
}
func abs(x int)int {
    if x < 0 {
        return -1 * x
    }
    return x
}

复杂度分析：
时间复杂度：O(n): 尾递归，每个节点只需要计算一次
空间复杂度：O(n): 递归的层次

总结：
1. 递归有两种方式：自顶向下，自底向上（尾递归）两个时间复杂度是不一样的
2. 对于函数的一些细节，比如返回值代表有含义的时候还是需要多多练习的，如本道题中的 return height(root) >= 0
3. 这道题的尾递归挺像后序遍历的，left right 根
