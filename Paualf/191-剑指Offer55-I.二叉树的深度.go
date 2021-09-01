输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

提示：
节点总数 <= 10000

1. Clarification:

后序遍历：
分解的思路：
leftDepth := helper(root.Left)
rightDepth := helper(root.Depth)

return max(leftDepth,rightDepth) + 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return helper(root)
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftDepth := helper(root.Left)
    rightDepth := helper(root.Right)

    return max(leftDepth,rightDepth) + 1
}

func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}

层序遍历，将每一层一层的弹出，每一层怎么表示呢？使用size去表示，这个也是之前一直想不到的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queueNode := []*TreeNode{root}
    depth := 0

    for len(queueNode) > 0 {
        depth++
        size := len(queueNode)

        for size > 0 {
            node := queueNode[0]
            queueNode = queueNode[1:]
            if node.Left != nil {
                queueNode = append(queueNode, node.Left)
            }

            if node.Right != nil {
                queueNode = append(queueNode, node.Right)
            }

            size--
        }
    }

    return depth
}

2. 看题解：


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 递归栈调用空间


总结：
4.1: 以前遇到过好几次写后序遍历都没有写出来，这次终于写出来了。。。
