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
*/

1. Clearfication:
高度平衡二叉树定义：
每个节点的左右两个子树的高度差的绝对值不超过1

方法1:
自顶向下 计算每个节点的高度，然后计算它的左右子树高度，然后进行计算，中间有重复计算，时间复杂度比较高
height 函数
func height(root *TreeNode) int {
        if root == nil {
            return 0
    }
    left := height(root.Left)
    right := height(root.Right)
    return max(left,right) + 1
}
func max(x,y int)int {
     if x > y {
        return x
   }
   return y
}

方法2: 
自底向上：
类似二叉树后序遍历的思想，左右根，然后计算从底计算每个二叉树的高度

方法1 Code:
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
     left := height(root.Left)
     right := height(root.Right)
     return abs(left,right) <=1 && isBalanced(root.Left) && isBalanced(root.Right) 
}
func abs(x ,y int)int {
    if x - y < 0 {
        return y - x
    }
    return x - y
}
func height(root *TreeNode) int {
        if root == nil {
            return 0
    }
    left := height(root.Left)
    right := height(root.Right)
    return max(left,right) + 1
}
func max(x,y int)int {
     if x > y {
        return x
   }
   return y
}

方法2，思想是知道了，换成自己写的时候还是卡壳了
看了题解：
func isBalanced(root *TreeNode) bool {
         return height(root) >= 0
}
func height(root *TreeNode) int {
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
func max(x,y int)int {
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
时间复杂度：第一种方法 ：自顶向下O(nlogn) 最坏O(n*n),n个节点，然后计算高度的话，平均是log(n),最坏情况下成一个链表时间复杂度为O(n),所以最坏时间复杂度是O(n*n)
第二种方法：时间复杂度是O(n) ，每个节点遍历一次

空间复杂度： 第一种方法：O(n): 递归栈空间,第二种方法：O(n) : 递归栈空间

总结：
1. 第一种方法写出来了，还是有点卡壳，思路没有特别的清晰

2. 第二种方法自己没有想出来，卡在了返回值 height(root) >=0 上，同时使用 -1 去标记不是对称二叉树

3. 第二种方法是变形的二叉树后序遍历
