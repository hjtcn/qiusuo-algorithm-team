/*
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
 
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
 
进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
*/

1. Clearfication:
镜像对称：
节点的左子树和节点的右子树相等，如果不相等 return false
找重复子问题: 
DFS 你会怎么写？
发现自己没有思路。。就去看了题解
核心思路是：使用两个指针进行遍历，一个向左一个向右
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
     return check(root, root)
}
func check(p,q *TreeNode) bool {
        if p == nil && q == nil {
            return true
    }
    if p == nil || q == nil {
            return false
    }
    
    return p.Val == q.Val && check(p.Left,q.Right) && check(p.Right, q.Left)
}

BFS:
核心思路是每次从队列中取出两个节点u,v 放入队列中u.Left,v.Right 和 u.Right 和 v.Left,然后进行比较
func isSymmetric(root *TreeNode)bool {
        u,v := root,root
    q := []*TreeNode{}
    q = append(q,u)
    q = append(q,v)
    for len(q) > 0 {
          u,v = q[0],q[1]
        q = q[2:]
        if u == nil && v == nil {
                continue
        }
        
        if u == nil || v == nil {
                return false
        }
        if u.Val != v.Val {
            return false
        }
        q = append(q,u.Left)
        q = append(q,v.Right)
        
        q = append(q,u.Right)
        q = append(q,v.Left)
    }
    return true
}
如果让我再写一遍这个题，我比较容易出错的地方在于 if 条件的判断:
 if u == nil && v == nil {
                continue
        }
        
        if u == nil || v == nil {
                return false
        }
        if u.Val != v.Val {
            return false
        }

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) DFS递归调用栈空间 或者是BFS开辟队列进行存储节点

总结： 
1. 没有想到使用两个指针遍历二叉树

2. 对二叉树的遍历还是没有真正的掌握，没有想清楚它的遍历
