//给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
//
//
//
// 示例 ：
//
// 输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]
//
//       5
//      / \
//    3    6
//   / \    \
//  2   4    8
// /        / \
//1        7   9
//
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
//            \
//             7
//              \
//               8
//                \
//                 9
//
//
//
// 提示：
//
//
// 给定树中的结点数介于 1 和 100 之间。
// 每个结点都有一个从 0 到 1000 范围内的唯一整数值。
//
// Related Topics 树 深度优先搜索 递归
// 👍 131 👎 0
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
func increasingBSTOld(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var nums = make([]int, 0)
	inOrderTraversal(root, &nums)

	new_tree := &TreeNode{Val:0}

	tmp := new_tree
	for _, val := range nums {
		next := &TreeNode{Val:val}
		tmp.Right = next
		tmp = tmp.Right
	}

	return new_tree.Right
}

func inOrderTraversal(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrderTraversal(root.Right, nums)
}

/**
中序遍历+重新构建树
先进行遍历将节点元素存入切片中，然后一次遍历切边，重新构造一个新树
复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n) 需要将所以节点进行存储
*/

var cur *TreeNode
func increasingBST(root *TreeNode) *TreeNode {
	new_tree := &TreeNode{Val:0}
	cur = new_tree
	inOrder(root)

	return new_tree.Right
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	node.Left = nil
	cur.Right = node
	cur = cur.Right
	inOrder(node.Right)
}
/**
中序遍历+更改树的连接方式
使用中序遍历，将树之间的节点进行重新连接，不用使用额外的空间存储所有节点
当遍历到一个节点时，左孩子设为空，其本身作为上一个遍历到节点的右孩子
复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(H) 递归使用栈空间，最多为树的高度H
*/

//leetcode submit region end(Prohibit modification and deletion)
