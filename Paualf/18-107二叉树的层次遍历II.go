/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：
[
  [15,7],
  [9,20],
  [3]
]
*/

使用队列进行层次遍历：
levelOrder 的时候从顶向下的时候

func levelOrder(root *TreeNode) [][]int {
    ret := make([][]int, 0)
    if root == nil {
        return ret
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        tmp := make([]int, 0)
        
        for size > 0 {
            size--
            node := queue[0]
            queue = queue[1:]
            tmp = append(tmp,node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ret = append(ret, tmp)
    }
    return ret
}

结果数组为从顶向下
题目要求是从底向上，所以我们需要得出结果后，需要reverse 一下数组

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    ret := make([][]int, 0)
    if root == nil {
        return ret
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        tmp := make([]int, 0)
        
        for size > 0 {
            size--
            node := queue[0]
            queue = queue[1:]
            tmp = append(tmp,node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ret = append(ret, tmp)
    }
    ret = reverse(ret)
    return ret
}
func reverse(ret [][]int)[][]int {
     for i, j := 0, len(ret)-1; i <= j; i, j = i+1, j-1 {
        ret[i],ret[j] = ret[j],ret[i]
    }
    return ret
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 使用队列存储节点信息

看了题解，dfs可以这样写：
还是蛮厉害的，遍历到nil的时候返回，将结果放到数组中

func levelOrderBottom(root *TreeNode) [][]int {
    result := [][]int{}
    helper([]*TreeNode{root}, &result)
    return result[1:]
}
func helper(level []*TreeNode, result *[][]int) {
    if len(level) == 0 {
        return
    }
    nextLevel := []*TreeNode{}
    for _, v := range level {
        if v != nil {
            nextLevel = append(nextLevel, v.Left)
            nextLevel = append(nextLevel, v.Right)
        }
    }
    helper(nextLevel, result)
    list := []int{}
    for _, v := range level {
        if v != nil {
            list = append(list, v.Val)
        }
    }
    *result = append(*result, list)
}

还有可以不使用 reverse 的解法
根据temp新建二维数组，然后将ret append 到新建的结果数组中

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var ret [][]int
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        temp := make([]int, 0, len(queue))
        for n := len(queue); n > 0; n-- {
            node := queue[0]
            queue = queue[1:]
            temp = append(temp, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ret = append([][]int{temp}, ret...)
    }
    return ret
}

总结：
1. 看题解，去看别人的思路和解法，语法是固定的，但是人的思维和写出来的代码是有想象力的，要去多看多学习
2. 语法是规定的，使用语法写出的优秀的项目是灵活的，是有思想的，有灵魂的
