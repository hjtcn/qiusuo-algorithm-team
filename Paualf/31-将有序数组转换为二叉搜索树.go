/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
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
一开始看这道题，感觉和上周五的题目比较像
总体思路是：选取中间节点为根节点，根节点左边为左子树，根节点右边为右子树
细节： 
终止条件，root == nil 的时候 return
下标计算使用

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
    mid := len(nums) / 2 
    
    root := &TreeNode{
        Val:nums[mid],
        Left:nil,
        Right:nil,
    }
    root.Left = sortedArrayToBST(nums[0:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}

更加简洁的写法

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
        Left: sortedArrayToBST(nums[:len(nums)/2]),
        Right: sortedArrayToBST(nums[len(nums)/2+1:]),
    }
}

官方理解中使用到了helper 函数进行解决：

func sortedArrayToBST(nums []int) *TreeNode {
    return helper(nums, 0,len(nums) -1)
}
func helper(nums []int,left,right int)*TreeNode {
    if left > right {
        return nil
    }
    mid := (left + right) / 2
    root := &TreeNode{Val:nums[mid]}
    root.Left = helper(nums,left,mid - 1)
    root.Right = helper(nums,mid + 1, right)
    return root
}

通常有些时候需要辅助变量的时候，我们使用helper函数更容易处理和解决
思路：构建二叉搜索树的过程感觉还是利用了二叉搜索树的特性，中序遍历为有序数组，我们想要左右两个子树的高度差的绝对值不超过 1，我们可以使用中间节点作为根，然后中间节点左边作为左子树，右边作为右子树

总结：

1. 需要了解二叉搜索树的特性

2. 使用递归进行构造，首先要有思路，有思路以后要注意细节变量，参数，还有对应的返回值
