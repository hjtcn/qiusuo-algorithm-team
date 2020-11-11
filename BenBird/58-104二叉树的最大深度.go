//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索
// 👍 734 👎 0
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepthOld(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/**
递归、DFS

利用递归，分别获取二叉树的左右节点的深度，取深度最大即可

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(height)
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0

	for len(queue) > 0 {
		sz := len(queue)
		for ; sz > 0; sz-- {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans++
	}

	return ans
}

/**
BFS(广度优先搜索)

逐层遍历二叉树节点，遍历完一层，深度+1，直到遍历完最后一层，二叉树最大深度即可得出

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */

//leetcode submit region end(Prohibit modification and deletion)

/** 测试代码 */

/**
构建二叉树
*/
func buildBinaryTree(params_slice []int, key int) *TreeNode {
	new_node := &TreeNode{params_slice[key], nil, nil}
	slice_len := len(params_slice)

	if key < slice_len && (2 * key + 1) < slice_len {
		new_node.Left = buildBinaryTree(params_slice, 2 * key + 1)
	}

	if key < slice_len && (2 * key + 2) < slice_len {
		new_node.Right = buildBinaryTree(params_slice, 2 * key + 2)
	}

	return new_node
}

func main() {
	params_slice := []int{3,9,20,0,0,15,7}
	temp_tree := buildBinaryTree(params_slice, 0)
	res := maxDepth(temp_tree)

	fmt.Println(res)
}

/**
总结：
递归实现：代码简练，易于理解，BFS：代码量较多，理解起来相对比较费劲，还是万能的画图理解好使
其实本道题不难，但是做题时还是卡着了，卡点在如何构建一个二叉树，对在本地测试有执念，所以不断尝试去构建二叉树，方便在本地进行测试
不过，自己构建二叉树，然后对解题代码进行测试，对二叉树的学习和理解还是有一定帮助，就是耗点时间
 */
