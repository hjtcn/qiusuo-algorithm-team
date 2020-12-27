//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例:
//
// 输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 381 👎 0
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

/**
BFS
*/
func rightSideViewBFS(root *TreeNode) []int {
	queue := []*TreeNode{}
	res := []int{}

	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if i == length - 1 {
				res = append(res, node.Val)
			}
		}

	}

	return res
}

/**
DFS
*/
var res []int

func rightSideView(root *TreeNode) []int {
	res = []int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if (depth == len(res)) {
		res = append(res, root.Val)
	}

	depth++
	dfs(root.Right, depth)
	dfs(root.Left, depth)
}
//leetcode submit region end(Prohibit modification and deletion)

/**
题解：
BFS：广度优先遍历，逐层遍历，保存每层最后一个节点元素即可
复杂度分析：
	时间复杂度：O(n) 每个节点最多进出队列一次
	空间复杂度：O(n) 每个节点最多进出队列一次, 最多占用O(n)
DFS：
复杂度分析：深度优先遍历，先遍历根节点，在遍历右子树，最后遍历左子树，这样，保存每层最先遍历的节点元素即可
	时间复杂度：O(n) 最坏情况下访问n个节点
	空间复杂度：O(n) 最坏站内会包含接近树高度的节点数量，占用O(n)空间
*/
