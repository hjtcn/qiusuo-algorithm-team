给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。

1. Clarification:
将二叉搜索树的节点转换为累加树（Greater Sum Tree), 使每个节点node 的新值等于原树中大于或等于 node.val 的值之和。

2. 看题解：
看完题解以后发现自己把题目理解错了
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    sum := 0 
    var dfs func(*TreeNode)

    dfs = func(node *TreeNode) {
        if node != nil {
            dfs(node.Right)
            sum += node.Val
            node.Val = sum
            dfs(node.Left)
        }
    }

    dfs(root)
    return root
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
Q :这道题为什么不会做呢？
A：自己把题目理解错了，转换为累加树以后，累加树并不是二叉搜索树了

Q: 那么为什么会理解错呢？
A: 是不是自己没有真正的理解题目，而是带着自己的认识去理解的呢
