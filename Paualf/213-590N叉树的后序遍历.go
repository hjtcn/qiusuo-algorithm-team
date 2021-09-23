给定一个 N 叉树，返回其节点值的 后序遍历 。
N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

进阶：
递归法很简单，你可以使用迭代法完成此题吗?

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[5,6,3,2,4,1]

示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
 
提示：
N 叉树的高度小于或等于 1000
节点总数在范围 [0, 10^4] 内

1. Clarification:

左右根

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}

func helper(root *Node,ret *[]int) {
    // terminator
    if root == nil {
        return
    }
    // process current logic
    for i := 0;i < len(root.Children);i++ {
        helper(root.Children[i],ret)
    }

    *ret = append(*ret, root.Val)
}
对比二叉树的遍历
helper(root.Left,ret)
helper(root.Right,ret)

=>

for i := 0;i < len(root.Children);i++ {
    helper(root.Children[i], ret)
}
前序遍历和中序遍历的迭代都是先把左节点放进去
后序遍历的迭代呢？没有啥思路

2. 看题解：
后序遍历的迭代二叉树是怎么做的呢？

func postorderTraversal(root *TreeNode) []int {
    // 通过lastVisit标识右子节点是否已经弹出
    if root == nil {
        return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    var lastVisit *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 这里先看看，先不弹出
        node:= stack[len(stack)-1]
        // 根节点必须在右节点弹出之后，再弹出
        if node.Right == nil || node.Right == lastVisit {
            stack = stack[:len(stack)-1] // pop
            result = append(result, node.Val)
            // 标记当前这个节点已经弹出过
            lastVisit = node
        } else {
            root = node.Right
        }
    }
    return result
}
看了下二叉树的后序迭代遍历，对比了下，发现难理解的地方在标记 lastVisit 的地方
因为，如果节点的右边元素有值且未访问过的的话，就不能弹出，所以 lastVisit 是标记这个地方的处理。

虽然是一个小点，但是如果不是很好的理解还是不容易写出来的。

N叉树的遍历：后进先出
func postorder(root *Node) []int {
	if root==nil{
		return nil
	}
	stack := make([]*Node,0)
	stack = append(stack,root)
	res := make([]int,0)

	for len(stack)>0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val},res...)

		for i:=0;i<len(node.Children);i++{
			stack = append(stack,node.Children[i])
		}
        
	}
	return res
}
还看到先放进去，然后再交换元素的
func postorder(root *Node) []int {
	var res = []int{}
    if root == nil{
        return res
    }
	var stack = []*Node{root}
	for 0 < len(stack) {
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		l := len(root.Children)
		for i := 0; i < l; i++ {
			stack = append(stack, root.Children[i]) //入栈
		}
	}

	l := len(res) - 1
	for i := 0; i < l/2+1; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}
	return res
}

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

总结：
4.1: 文字里面记载着智慧的结晶和一些思考
