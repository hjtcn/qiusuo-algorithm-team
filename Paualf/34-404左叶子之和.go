/*
计算给定二叉树的所有左叶子之和。
示例：
    3
   / \
  9  20
    /  \
   15   7
在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/


Clearfication:
如何找到左叶子呢？有点疑惑
没想出来。。。
看了题解：
我们还是需要遍历每个节点，然后判断该节点是不是左子节点
那么左子节点的定义是什么呢：
node.Left 为 node，然后该 node 的Left 和 Right 都为nil,这个时候该节点为左子节点
然后利用递归，求解重复子问题：

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isLeafNode(node *TreeNode)bool {
    return node.Left == nil && node.Right == nil
}
func dfs(node *TreeNode)(ans int) {
    if node.Left != nil {
        if isLeafNode(node.Left) {
            ans += node.Left.Val
        }else {
            ans += dfs(node.Left)
        }
    }
    if node.Right != nil && !isLeafNode(node.Right){
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

mine thinking  it is wrong code
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isLeafNode(node *TreeNode)bool {
    return node.Left == nil && node.Right == nil
}
func dfs(node *TreeNode,ans *int) {
    if node.Left != nil {
        if isLeafNode(node.Left) {
            *ans += node.Left.Val
        }else {
            *ans += dfs(node.Left,ans)
        }
    }
    if node.Right != nil && !isLeafNode(node.Right){
        *ans += dfs(node.Right,ans)
    }
}
func sumOfLeftLeaves(root *TreeNode) int {
     ans := 0
     if root == nil {
         return 0
     }
     return dfs(root,&ans)
}

看到题解里面有这样的题解，还是比较清奇的
func sumOfLeftLeaves(root *TreeNode)int {
     res := 0
   visit(root, false, &res)
   return res
}
func visit(root *TreeNode,isLeft bool,res *int) {
        if root == nil {
            return
    }
    
    if root.Left == nil && root.Right == nil && isLeft {
            *res += root.Val
    }
    
    visit(root.Left, true, res)
    visit(root.Right, false, res)
}

这个代码感觉也挺好的,返回int的时候返回了0
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left != nil && isLeaf(root.Left) {
        return root.Left.Val + sumOfLeftLeaves(root.Right)
    }
    return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
func isLeaf(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
        return true
    }
    return false
}

BFS:
知道上面如何判断左叶子节点了，bfs就知道怎么做了,加一下判断是否为左子节点

func isLeafNode(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
                return true
    }
    return false
}
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    queue := []*TreeNode{}
    
    if root == nil {
         return sum
    }
    
    queue = append(queue, root)
    
    for len(queue) > 0 {
          node := queue[0]
        queue = queue[1:]
        
        if node.Left != nil {
                queue = append(queue, node.Left)
            if isLeafNode(node.Left) {
                    sum += node.Left.Val
            }
        }
        
        if node.Right != nil {
                queue = append(queue,node.Right)
        }
    }
    
    return sum
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归栈或者开辟空间存储

总结：
1. 树的题目变化比较多，多练多想

2. 找重复子问题

3. 要多看代码，官方题解写的代码有的时候并不咋的
