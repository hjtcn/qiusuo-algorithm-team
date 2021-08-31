给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

1. Clarification:

层序遍历：
使用队列，一层一层的将元素放到一维数组中，然后将一维数组放到二维数组中
使用dfs，记录层数，然后按照层数将数据push到对应的二维数组中

append 的时候有点问题
func levelOrder(root *TreeNode) [][]int {
    ret := [][]int{}

    if root == nil {
        return ret
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)
        tmp := []int{}

        for size > 0 {
            size--
            node := queue[0]
            queue = queue[1:]
            tmp = append(tmp,node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        ret = append(ret,tmp)
    }

    return ret
}
dfs的代码
func helper(root *TreeNode, ret *[][]int,level int) {
    if root == nil {
        return
    }
    
    *ret[level] = append(*ret, root.Val)
    helper(root.Left, ret,level + 1)
    helper(root.Right, ret,level + 1)
}

2.看题解：

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}
func dfs(root *TreeNode, level int) {
	if root != nil {
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}

func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    helper(root, 0, &ans)
    return ans
}

func helper(root *TreeNode, level int, ans *[][]int) {
    if root == nil {
        return
    }
    
    if len(*ans) <= level {
        *ans = append(*ans, []int{})    
    }
    
    (*ans)[level] = append((*ans)[level], root.Val)
    
    level++
    helper(root.Left, level, ans)
    helper(root.Right, level, ans)
}

上面的代码里面  if len(*ans) <= level 还是解决了我的疑惑的


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n): 队列空间or递归调用栈空间


4.总结：
4.1: 好像真的是：只要你做事情就会遇到问题，重要的是你遇到了问题怎么办，怎么思考，怎么解决

4.2: 遇到问题的时候，可以慢一点思考一下，不要逃避，想办法去解决
