计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

1. Clarification:

Q：什么是左叶子？怎么定义左叶子呢？
A: 没有想清楚什么是左叶子

2. 看题解：
判断左叶子节点的定义：
    node.Left 是否为左叶子节点，node.Left 的左节点和右节点都为nil
func isLeafNode(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
    if node.Left != nil {
        if isLeafNode(node.Left) {
            ans += node.Left.Val
        } else {
            ans += dfs(node.Left)
        }
    }
    if node.Right != nil && !isLeafNode(node.Right) {
        ans += dfs(node.Right)
    }
    return
}

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root)
}

func isLeafNode(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
    if root == nil {
        return
    }
    q := []*TreeNode{root}
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node.Left != nil {
            if isLeafNode(node.Left) {
                ans += node.Left.Val
            } else {
                q = append(q, node.Left)
            }
        }
        if node.Right != nil && !isLeafNode(node.Right) {
            q = append(q, node.Right)
        }
    }
    return
}

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

总结：
4.1: 知识点知道，但是遇到问题能不能灵活运用和想起来还是需要练习、总结、思考、反思，想一下为什么
