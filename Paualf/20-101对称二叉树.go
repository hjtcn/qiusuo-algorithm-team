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

递归：
1. 找重复性
不为空的时候满足
为镜像对称
2. 遍历树结构
思路不是很清晰，没有想出来
想不清楚的点在于代码如何实现和组织 node 与 相领的node 的 子树比较，例子中也就是 
左节点2和右点2，如何遍历到，然后进行比较，卡壳的地方 

看了题解：
看到用了两个指针还是很巧妙的
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
    return p.Val == q.Val && check(p.Left,q.Right) && check(p.Right,q.Left)
}

复杂度分析：
时间复杂度：O(n)
空间复杂度：O(height) 递归使用到栈，所有需要有使用到空间的
迭代：
 还是没想出来，其实思路和递归是一样的，所以就去看了题解：

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
      u,v := root,root
      q := []*TreeNode{}
      q = append(q, u)
      q = append(q, v)
      for len(q) > 0 {
          node1 := q[0]
          node2 := q[1]
          q = q[2:]
          if node1 == nil && node2 == nil {
              continue
          }
          if node1 == nil || node2 == nil {
              return false
          }
          if node1.Val != node2.Val {
              return false
          }
          q = append(q,node1.Left)
          q = append(q,node2.Right)
          q = append(q,node1.Right)
          q = append(q,node2.Left)
      }
      return true
}

复杂度分析:
时间复杂度：O(n) ：遍历树的节点个数
空间复杂度：O(n)： 使用队列存储节点

总结：
1. 题目是做过，然后重新做还是没有思路，说明还是不会
2. 做题是一方面，同时要找方法，找思路，做总结
