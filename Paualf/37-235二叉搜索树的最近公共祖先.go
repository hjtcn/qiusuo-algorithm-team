/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6 
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 
说明:
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
*/

Clearfication:
 这道题目做过，发现思路还是比较少，怎么找他们的最近公共祖先，如何定义这个问题是比较关键的，递归思路是比较难想的
迭代的话，可以通过遍历保存它的父亲节点的路径，然后在p的路径中找到q的路径，它们公共的部分

预估复杂度：
时间复杂度：O(n) 遍历每个节点需要的时间复杂度
空间复杂度：O(n) 递归开辟的栈空间或者队列开辟存储的节点元素

Coding:
尽可能的分析：
DFS:
找重复子问题：
flson frson
BFS:
如何保存节点路径
最后如何从一条路径中找到另外一条路径
看了题解以后对比我们做过的题，找路径的时候，可以利用二叉搜索树，更快的构建路径

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func getPath(root, target *TreeNode)[]*TreeNode {
    path := []*TreeNode{}
    
    node := root
    for node != target {
        path = append(path, node)
        if target.Val < node.Val {
            node = node.Left
        }else{
            node = node.Right
        }
    }
    path = append(path,node)
    return path
}
func lowestCommonAncestor(root,p,q *TreeNode) *TreeNode {
    ancestor := &TreeNode{}
    pathP := getPath(root, p)
    pathQ := getPath(root, q)
    for i := 0;i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i];i++ {
        ancestor = pathP[i]
    }
    return ancestor
}

func getPath(root, target *TreeNode)[]*TreeNode {
    path := []*TreeNode{}
    
    node := root
    // 这个地方判断的是节点是否相等，所以才有循环外的 path = append(path,node)
    for node != target {
        path = append(path, node)
        if target.Val < node.Val {
            node = node.Left
        }else{
            node = node.Right
        }
    }
    path = append(path,node)
    return path
}

一开始没有想明白为什么循环外还要 path = append(path, node)，后来发现是自己没有注意到 for 循环的判断条件 node != target 判断的是节点是否相等,而不是节点的值是否相等，说明自己还是没有明白这个地方的题意 =》 是找节点是否相等
看了中等难度的BFS:
发现如果自己再写一遍还是写不出来

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

DFS:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode {
    ancestor := root
    
    for {
        if p.Val < ancestor.Val && q.Val < ancestor.Val {
            ancestor = ancestor.Left
        }else if p.Val > ancestor.Val && q.Val > ancestor.Val {
            ancestor = ancestor.Right
        }else {
            break
        }
    }
    
    return ancestor
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归栈空间 or 使用到的队列开辟使用的空间

总结：
1. 定义问题，将问题分析清楚这种能力需要不断练习

2. 对比类似题目，通过对比进行分析，找相同和不同点
