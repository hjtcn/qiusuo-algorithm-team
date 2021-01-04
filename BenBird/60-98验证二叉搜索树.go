//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 递归
// 👍 870 👎 0
package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBSTOld(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool  {
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
//leetcode submit region end(Prohibit modification and deletion)

/**
题解：
1. 递归：根据二叉搜索树特点，当前节点必定大于左子树所有节点，小于右子树上的所有节点，左右子树也为二叉搜索树，所以使用递归检查左右子树是否
		满足二叉搜索树特点
	复杂度分析：
		时间复杂度：O(n)因每个节点只访问一次，n个节点，时间复杂度为O(n)
		空间复杂度：O(n), n为二叉树节点, 递归时，需要为每一层递归分配栈空间，空间复杂度取决于树的高度，所以最坏情况二叉树为一条链，树的高度为n
				  ,即最坏情况下的空间复杂度为O(n)
2. 中序遍历：根据中序遍历的特点，中序遍历的序列是升序的，所以只需要判断当前值是否大于前一个值，这段代码不易懂，需结合实际二叉树进行理解
	复杂度分析：
		时间复杂度：O(n) 每个节点只访问一次
		空间复杂度：O(n) 栈内最多存储n个二叉树节点，额外空间最多是O(n)
*/
