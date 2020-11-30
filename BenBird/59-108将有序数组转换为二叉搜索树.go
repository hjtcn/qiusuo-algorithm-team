//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 树 深度优先搜索
// 👍 645 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return helperLeft(nums, 0, len(nums) - 1)
}

func helperLeft(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2	//始终以中间偏左为根节点
	//mid := (left + right + 1) / 2	//始终以中间偏右边为根节点
	//rand.Seed(time.Now().UnixNano())
	//mid := (left + right + rand.Intn(2)) / 2	//随机取中间偏左或偏右为根节点
	root := &TreeNode{Val: nums[mid]}
	root.Left = helperLeft(nums, left, mid - 1)
	root.Right = helperLeft(nums, mid + 1, right)

	return root
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{}
	res := sortedArrayToBST(nums)
	fmt.Println(res)
}

/**
题解：

由题目可知，所给数组为有序数组，而二叉树的中序遍历也是升序的，所以本题，可以理解为给定二叉树的中序遍历，还原二叉树，
但条件是还原的二叉树需要是高度平衡，即左右两子树的高度差不超过1

所以本题解决的思路主要是解决根节点选取，使用递归

所以根节点尽量选取有序数组中间位置的元素
若数组元素总数为奇数，直接选取中间位置的数字即可
若数组元素总数为偶数的话，可以有选择的选取中间偏左或偏右的元素为根节点

复杂度分析：
	时间复杂度：O(N) 数组长度为N的数组，每个元素都会访问一次
	空间复杂度：O(logN)	N为数组的长度，空间复杂度不考虑返回值，主要取决于递归栈的深度，即O(logN)
 */
