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

p,q 有三种不同情况
1. p,q 在root 的节点的两边，root为最近公共祖先
2. p = root || q = root 返回root
3. p = 
第一种情况如 6，0 他们的最近公共祖先为3
第二种情况，2，4 他们的最近公共祖先为 22
自己还是分析的不是很清楚，还是没有将3种情况分析清楚
看了一下题解，感觉只有两种情况

1. 两个节点在x 节点的两边，这个时候 (flson && frson)
2. (x == p || x ==q) && (flson || frson)

func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == p.Val || root.Val == q.Val {
        return root
    }
    
    left := lowestCommonAncestor(root.Left,p,q)
    right := lowestCommonAncestor(root.Right,p,q)
    
    if left != nil && right != nil {
        return root
    }
    
    if left == nil {
        return right
    }
    return left
}

1. 用哈希表记录每个节点的父节点指针
2. 从p节点开始不断忘它的祖先移动，并用数据结构记录已经访问过的祖先节点
3. 从q节点不断往它的主线移动，如果有祖先访问过，即意味着这是p和q的深度最深的公共祖先

func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode {
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

利用二叉搜索树的性质：
func getPath(root,target *TreeNode) []*TreeNode {
    path := []*TreeNode{}
    for root != target {
        path = append(path, root)
        if target.Val < root.Val {
            root = root.Left
        }else {
            root = root.Right
        }
    }
    path = append(path, root)
    return path
}

func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode {
    pathP := getPath(root,p)
    pathQ := getPath(root,q)
    
    for i := 0;i < len(pathP) && i < len(pathQ) && pathP[i] ==pathQ[i];i++ {
        ancestor = pathP[i]
    }
    
    return ancestor
}

func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode {
    ancestor := root
    for {
        if p.Val < ancesor.Val && q.Val < ancestor.Val {
            ancestor = ancestor.Left
        }else if p.Val > ancestor.Val && q.Val > ancestor.Val {
            ancestor = ancestor.Right
        }else {
            break;
        }
    }
    return ancestor
}

复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) or O(1)

总结：
1. 不知道咋回事，可能是最近有点忙或者其他情况，感觉挺累的，一点点分析加抄题解，将题目写完了。。。
