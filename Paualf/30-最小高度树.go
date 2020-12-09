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

Clearfication:
建立高度最小的平衡二叉树：
脑袋里又问题产生了：
a. 怎么构建平衡二叉树
b. 构建的平衡二叉树什么情况下是最小的
没有思路哈：
看了题解：

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    if len(nums) == 1 {
        return &TreeNode{Val:nums[0]}
    }
    
    middle := len(nums) / 2
    head := &TreeNode{Val:nums[middle]}
    head.Left = sortedArrayToBST(nums[:middle])
    head.Right = sortedArrayToBST(nums[middle+1:])
    return head
}

看到一个更简洁的：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums) / 2],
        Left: sortedArrayToBST(nums[:len(nums) / 2]),
        Right: sortedArrayToBST(nums[len(nums)/2 + 1:]),
    }
}

复杂度分析：
时间复杂度：O(n) ：遍历每个节点
空间复杂度：O(height(n)）: 递归使用到的栈空间

复盘：
1. 没看到是顺序数组，可能是昨天的bug讨论以至于今天早上没有看清楚题目
2. 高度最小的时候，要保持平衡，这个时候需要左右子树的节点个数尽可能接近
递增数组：利用二叉搜索树的特性，高度最小，左右子树节点个数相同，取中间节点作为树根，将[左区间：中间节点]  作为 left 将 [中间节点+1:] 作为right，递归该过程创建二叉树

总结：
1. 要了解二叉搜索树的特性

2. 树的问题为什么经常会用递归解决，因为树的问题，大部分都是解决重复的子问题，将子问题解决完，然后大问题也就解决了
