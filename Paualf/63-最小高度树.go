/*
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
          0 
         / \ 
       -3   9 
       /   / 
     -10  5 
*/

1. Clearfication:
建立一个高度最小的二叉搜索树
以中间节点为根节点，左边节点为根节点的左子树，右边节点为根节点的右子树，递归该过程创建二叉树

if len(nums) <= 0 {
}
mid := len(nums) / 2
tree = &TreeNode{
nums[mid],
Left:nil,
Right:nil,
}
tree.Left = build(nums[0:mid])
tree.Right = build(nums[mid+1:])
return tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
         if len(nums) <= 0 {
            return nil
     }
     mid := len(nums) / 2
     tree := TreeNode{
            Val:nums[mid],
        Left:nil,
        Right:nil,
     }
     
     tree.Left = build(nums[0:mid])
     tree.Right = build(nums[mid+1:])
     return tree
}
上面代码思路是对的，初始化结构体的时候报错了，应该是&TreeNode,同时没有build 这个函数的，应该是sortedArrayToBST
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
         if len(nums) <= 0 {
            return nil
     }
     mid := len(nums) / 2
     tree := &TreeNode{
            Val:nums[mid],
        Left:nil,
        Right:nil,
     }
     
     tree.Left = sortedArrayToBST(nums[0:mid])
     tree.Right = sortedArrayToBST(nums[mid+1:])
     return tree
}

复杂度分析：
时间复杂度度：O(n)： 遍历每个节点
空间复杂度：O(n) ： 递归创建二叉树调用栈空间

总结：
不要害怕，分析清楚，动手写
