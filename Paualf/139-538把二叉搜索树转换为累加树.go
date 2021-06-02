给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
提醒一下，二叉搜索树满足下列约束条件：
节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同
示例 1：
输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：
输入：root = [0,null,1]
输出：[1,null,1]
示例 3：
输入：root = [1,0,2]
输出：[3,3,2]
示例 4：
输入：root = [3,2,4,1]
输出：[7,9,4,10]
 
提示：
树中的节点数介于 0 和 104 之间。
每个节点的值介于 -104 和 104 之间。
树中的所有值 互不相同 。
给定的树为二叉搜索树。

1. Clearfication:
看了下题目，没有看懂是什么意思
累加树：使每个节点node的新值大于原树中或等于node.val 的值之和

2. Coding:
  3.看题解：
看了题解以后发现是倒着来的，一开始的36我没看懂是什么意思
自己试着写的时候发现不知道怎么去更新节点的值
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    var dfs func(*TreeNode)
    
    dfs = func(node *TreeNode) {
        dfs(node.Right)
        sum += node.Val
        node.Val = sum
        dfs(node.Left)
    }
    
    dfs(root)
    return root
}   

func converBST(root *TreeNode) *TreeNode {
    var sum = 0
    helper(root,&sum)
    return root
}
func helper(root *TreeNode,sum *int) {
    if root == nil {
        return
    }
    helper(root.Right, sum)
    *sum += root.Val
    root.Val = *sum
    helper(root.Left, sum)
}
自己没有想到迭代。。。
func convertBST(root *TreeNode) *TreeNode {
    stack := []*TreeNode{}
    node := root
    sum := 0
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Right
        }
        node = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        sum += node.Val
        node.Val = sum
        node = node.Left
    }
    return root
}

4. 复杂度分析：
时间复杂度：O(n) : 节点个数
空间复杂度：O(n) : 递归调用栈深度 or 迭代开辟栈深度

5. 总结：
5.1：二叉树遍历真的简单么？ =》 忘完了，换了个方式，都没有办法灵活应用，说明还是没有完全掌握
