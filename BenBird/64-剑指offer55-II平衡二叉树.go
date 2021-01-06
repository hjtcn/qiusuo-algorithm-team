//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回 true 。
//
//示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//        1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// 返回 false 。
//
//
//
// 限制：
//
//
// 1 <= 树的结点个数 <= 10000
//
//
// 注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/
//
//
// Related Topics 树 深度优先搜索
// 👍 100 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalancedOld(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := recur(root.Left)
	if left == -1 {
		return -1
	}

	right := recur(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left - right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

/**
后序遍历+剪枝
由下向上计算左右子树高度差，当任意左右子树高度差大于1时，即停止计算左右子树的高度差，返回结果
剪枝常见方式：可行性剪枝、排除等效冗余、最优性剪枝、顺序剪枝、记忆化

复杂度分析：
	时间复杂度：O(n)	最差情况，遍历所有节点
	空间复杂度：O(n)	递归调用使用栈空间
*/

func abs(number int) int {
	if number >= 0 {
		return number
	}

	return -number
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depth(root.Left) - depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}
/**
先序遍历+深度判断
由上到下递归对比左右子树的深度，二叉树底部节点重复遍历多次
复杂度分析：
	时间复杂度：O(NlogN) 满二叉树的层数为log2(N+1), 高度的复杂度为O(logN)，
		每层遍历的复杂度为O(n)，多以总时间复杂度为 O(NlogN)
	空间复杂度：O(n) 递归调用使用栈空间
*/

//leetcode submit region end(Prohibit modification and deletion)
