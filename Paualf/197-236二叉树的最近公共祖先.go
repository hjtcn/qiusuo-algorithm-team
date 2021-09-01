给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
示例 2：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
示例 3：

输入：root = [1,2], p = 1, q = 2
输出：1
 
提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。

Clarification:

需要使用二叉树的后序遍历

Q：怎么找p和q?

Q: 怎么分解子问题？
没什么头绪

Q：记得有记录父亲节点的方式：
5 -> 3
4 -> 2 -> 5 -> 3
这种怎么实现呢？
想了下还是不会写。。。。。。

看题解：
我们递归遍历整棵二叉树，定义f(x) 表示x节点的子树中是否包含p节点或q节点，如果包含为true,否则为false。那么：
(flson && frson) || (x == p || x ==q) && (flson||frson) 

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

看完题解以后感觉比较巧妙的地方在于：
  if root.Val == p.Val || root.Val == q.Val {
        return root
    }
这个地方的判断

方法二：
使用哈希表存储所有节点的父节点，然后我们就可以利用节点的父节点信息从p节点开始不断往上跳，并记录已经访问过的节点，再从q节点开始不断往上跳，如果碰到已经访问过的节点，那么这个节点就是我们要找的最近公共祖先
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
空间复杂度：O(n)

总结：
4.1: 知道，然后把这个知识点或者什么东西变成自己的东西，需要经过自己的思考和实践

4.2: 越来越感觉编程就像数学题，只有你把思路写清楚了，才能把问题给解决，而不是上来就动手写很多代码。

上来就写很多代码，无非是给自己徒增工作量罢了，你后面还是要删掉。给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
示例 2：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
示例 3：

输入：root = [1,2], p = 1, q = 2
输出：1
 
提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。

1. Clarification:

需要使用二叉树的后序遍历

Q：怎么找p和q?

Q: 怎么分解子问题？
没什么头绪

Q：记得有记录父亲节点的方式：
5 -> 3
4 -> 2 -> 5 -> 3
这种怎么实现呢？
想了下还是不会写。。。。。。

2. 看题解：
我们递归遍历整棵二叉树，定义f(x) 表示x节点的子树中是否包含p节点或q节点，如果包含为true,否则为false。那么：
(flson && frson) || (x == p || x ==q) && (flson||frson) 

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

看完题解以后感觉比较巧妙的地方在于：
  if root.Val == p.Val || root.Val == q.Val {
        return root
    }
这个地方的判断

方法二：
使用哈希表存储所有节点的父节点，然后我们就可以利用节点的父节点信息从p节点开始不断往上跳，并记录已经访问过的节点，再从q节点开始不断往上跳，如果碰到已经访问过的节点，那么这个节点就是我们要找的最近公共祖先
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

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 知道，然后把这个知识点或者什么东西变成自己的东西，需要经过自己的思考和实践

4.2: 越来越感觉编程就像数学题，只有你把思路写清楚了，才能把问题给解决，而不是上来就动手写很多代码。

上来就写很多代码，无非是给自己徒增工作量罢了，你后面还是要删掉。
