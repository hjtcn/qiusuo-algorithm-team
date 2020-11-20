/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
示例 1:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
*/
 
说明:
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
Clarfication:
这道题给我的第一感觉是找重复性：
卡在那里了
公共祖先：
root == p || root == q || root.left == p || root.left == q || root.right == p || root.right == q
return root
怎么判断它是最近公共祖先呢？
看了题解
这个分析公式才是正确的
(flson && frson) || ((x = p || x = q) && (flson || f rson))
还有题目中说了节点是一定存在树中的
递归代码虽然简单，但是理解起来还是有点复杂的

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == p.Val || root.Val == q.Val {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
}

复杂度分析：
时间复杂度：O(n) 遍历树节点个数
空间复杂度:  O(n) 递归使用到的栈空间O(height)

第二种思路：是将每个节点的父亲节点构造成一个map，然后将p节点经过的节点标记为true,再遍历q节点经历过的节点，找到第一个公共的返回
感觉是跳一跳的感觉，这个思路比递归容易理解

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent := map[int]*TreeNode{}
    visited := map[int]bool{}
    var dfs func(*TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }
        if r.Left != nil {
            parent[r.Left.Val] = r
            dfs(r.Left)
        }
        if r.Right != nil {
            parent[r.Right.Val] = r
            dfs(r.Right)
        }
    }
    dfs(root)
    for p != nil {
        visited[p.Val] = true
        p = parent[p.Val]
    }
    for q != nil {
        if visited[q.Val] {
            return q
        }
        q = parent[q.Val]
    }
    return nil
}

复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 构建hash关系，以及开辟访问过的数组空间还有递归树的栈空间

总结：
1. 递归还是找重复性，然后将重复性和自问题分析清楚，才能把代码写清楚

2. 写题目前一定要将题目分析清楚，这道题目就是属于分析不清楚不会做的类型
