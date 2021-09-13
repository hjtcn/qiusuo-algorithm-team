给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

示例 1：
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
        - 空数组，无子节点。
        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
            - 空数组，无子节点。
            - 只有一个元素，所以子节点是一个值为 1 的节点。
    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
        - 只有一个元素，所以子节点是一个值为 0 的节点。
        - 空数组，无子节点。

示例 2：
输入：nums = [3,2,1]
输出：[3,null,2,null,1]
 
提示：
1 <= nums.length <= 1000
0 <= nums[i] <= 1000
nums 中的所有整数 互不相同

Clarification:
找到数组中最大的值所在的索引位置，比如说是：k
 	

     node := &TreeNode{
 Val:nums[k]
        }
     node.Left = build(nums[0:k])
     node.Right = build(nums[k:])

返回值是一个啥，返回值是一个节点


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    k := findMaxIndex(nums)

    root := &TreeNode{
        Val:nums[k],
    }

    root.Left = constructMaximumBinaryTree(nums[:k])
    root.Right = constructMaximumBinaryTree(nums[k+1:])

    return root
}


func findMaxIndex(nums []int)int {
    k := 0
    max := nums[0]

    for i := 1;i < len(nums);i++ {
        if nums[i] > max {
            k = i
            max = nums[i]
        }
    }

    return k
}
一开始的时候
    root.Right = constructMaximumBinaryTree(nums[k+1:])
这里写成了 nums[k:]
程序就死循环了，超出内存限制了

看题解：

关注这一层怎么处理，这一层处理好了，数学归纳法，其他地方也就处理好了

复杂度分析：
时间复杂度：O(n * n) 递归创建时间复杂度为O(n),找出最大元素的时间复杂度也是O(n),所以是O(nxn)
空间复杂度：O(n)

总结：
4.1： 递归感觉也是交给我们一个道理，更加关注当下，不要想的太远。。。


