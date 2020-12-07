/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。
示例: 
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/

Clearfication:
路径上所有节点的值等于目标和
DFS，最后遍历到叶子节点，并且路径和为sum,则返回true,否则返回false
BFS，一层一层遍历到叶子节点，路径和为sum,则返回true,否则返回false

预估复杂度：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归使用到的栈空间 或者 BFS 使用遍历到的队列元素

Coding:
DFS：
这是我自己想的代码，返回值没有想清楚怎么返回，感觉还有地方怪怪的，如果root.Left == nil 的话，这样就不用遍历 root.Left 了，所以下面代码是有问题需要调整的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    // terminator
    if root == nil && sum == 0 {
        return true
    }
    // process current logic
    if root.Left != nil {
         hasPathSum(root.Left, sum - root.Val)
    }
   
    if root.Right != nil {
         hasPathSum(root.Right, sum - root.Val)
    }
}

修改
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    // terminator
    if root == nil {
        return false
    }
    // process current logic
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right,sum - root.Val)
}

BFS:
BFS 也写了一半，然后错了
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    queue := []*TreeNode{root}
    currentSum := 0
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        currentSum += node.Val
        if node.Left == nil && node.Right == nil && currentSum == sum {
            return true
        }
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return false
}

修改，还需要一个队列记录当前记录的和
func hasPathSum(root *TreeNode,sum int)bool {
    if root == nil {
        return false
    }
    queNode := []*TreeNode{}
    queVal := []int{}
    queNode = append(queNode, root)
    queVal = append(queVal, root.Val)
    
    for len(queNode) != 0 {
        now := queNode[0]
        queNode = queNode[1:]
        temp := queVal[0]
        queVal = queVal[1:]
        if now.Left == nil && now.Right == nil {
            if temp == sum {
                return true
            }
        }
        
        if now.Left != nil {
            queNode = append(queNode, now.Left)
            queVal = append(queVal, now.Left.Val + temp)
        }
        
        if now.Right != nil {
            queNode = append(queNode, now.Right)
            queVal = append(queVal, now.Right.Val + temp)
        }
    }
    return false
}

是从 ret = 0 一直加到 sum, 如果 ret == sum 然后返回，还是sum一直减，最后到0返回

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归栈空间 或者是 开辟存储队列空间

总结：
1. esay said than done,觉得题目很简单，自己还是有一些地方没有想清楚

2. 多想多练
