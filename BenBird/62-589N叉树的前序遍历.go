//给定一个 N 叉树，返回其节点值的前序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其前序遍历: [1,3,5,6,2,4]。
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树
// 👍 128 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package main

type Node struct {
	Val int
	Children []*Node
}

var res []int

/*递归*/
func preorderOld(root *Node) []int {
	res = []int{}
	dfsN(root)

	return res
}

func dfsN(root *Node)  {
	if root != nil {
		res = append(res, root.Val)
		for _, n := range root.Children {
			dfsN(n)
		}
	}
}

/*迭代*/
func preorder(root *Node) []int {
	stack := []*Node{}
	res := []int{}

	stack = append(stack, root)
	stack_len := len(stack)

	for len(stack) > 0 {
		stack_len = len(stack)
		temp_node := stack[stack_len - 1]
		stack = stack[:stack_len - 1]

		if temp_node == nil {
			continue
		}

		res = append(res, temp_node.Val)

		childred_len := len(temp_node.Children)
		if childred_len > 0 {
			for i := childred_len - 1; i >= 0; i-- {
				stack = append(stack, temp_node.Children[i])
			}
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)

/**
题解：
递归：先序遍历，首先遍历根节点，然后依次从左到右递归遍历根节点的子节点
	复杂度分析：
		时间复杂度：O(n)
		空间复杂度：O(n)

迭代：首先将根节点入栈，随后取出栈顶节点，进行存储，然后将栈顶节点的子节点按逆序入栈，然后再获取栈顶元素。。
	复杂度分析：
		时间复杂度：O(n)	所有节点都会入栈出栈一次
		空间复杂度：O(n)	最坏情况下，N叉树为两层，即栈中会存放n-1个节点元素，所以空间复杂度为O(n)
*/
