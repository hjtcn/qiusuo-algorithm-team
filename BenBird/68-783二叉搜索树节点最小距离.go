//给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
//
//
//
// 示例：
//
// 输入: root = [4,2,6,1,3,null,null]
//输出: 1
//解释:
//注意，root是树节点对象(TreeNode object)，而不是数组。
//
//给定的树 [4,2,6,1,3,null,null] 可表示为下图:
//
//          4
//        /   \
//      2      6
//     / \
//    1   3
//
//最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
//
//
//
// 注意：
//
//
// 二叉树的大小范围在 2 到 100。
// 二叉树总是有效的，每个节点的值都是整数，且不重复。
// 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
//相同
//
// Related Topics 树 递归
// 👍 109 👎 0
package main

import (
	"math"
	"sort"
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
func minDiffInBSTOld(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min_diff := math.MaxInt64
	curr := 0;
	minDiff(root, &min_diff, &curr)

	return min_diff
}

func minDiff(root *TreeNode, min_diff *int, curr *int) {
	if root == nil {
		return
	}

	minDiff(root.Left, min_diff, curr)
	if *curr != 0 {
		*min_diff = minNum(*min_diff, root.Val - *curr)
	}
	*curr = root.Val
	minDiff(root.Right, min_diff, curr)
}

func minNum (num_one, num_two int) int {
	if num_one < num_two {
		return num_one
	}

	return num_two
}
/**
中序遍历
二叉搜索树中序遍历的结果集是依次递增序列，所以在中序遍历时，进行两两相减并和当前最小值比较 取最小即可
中序遍历完成之后，即可获取任意两节点之差最小值

复杂度分析：
	时间复杂度：O(n) n为树中节点数
	空间复杂度：O(h) h为树的高度
*/

var node_set []int
func minDiffInBST(root *TreeNode) int {
	node_set = make([]int, 0)	//注意
	minDiffDfs(root)
	sort.Ints(node_set)

	min_diff := math.MaxInt64
	for i := 1; i < len(node_set); i++ {
		min_diff = minNum(min_diff, node_set[i] - node_set[i - 1])
	}

	return min_diff
}

func minDiffDfs(root *TreeNode) {
	if root == nil {
		return
	}

	node_set = append(node_set, root.Val)
	minDiffDfs(root.Left)
	minDiffDfs(root.Right)
}
/**
先序遍历+排序
通过先序遍历获取所有节点，然后进行排序
接着就进行依次遍历，进行两两相减，之差和最小值进行比较，获取最小值

复杂度分析：
	时间复杂度：O(nlogn)	递归+排序 n + nlogn 所以是O(nlogn)
	空间复杂度：O(n)

*/

//leetcode submit region end(Prohibit modification and deletion)
