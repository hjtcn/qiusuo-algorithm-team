/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

递归模版：
terminator ：终止条件 
process current logic ： 处理当前层逻辑
dirll down ： 下探到下一层
reverse current status ： 将当前层逻辑重置 
递归模版：
这道题：
1. terminator: 
if  root == nil {
    return 0
}
2. process current login & drill down 
depth = max(maxDepth(root.Left),maxDepth(root.Right)) + 1
3. no need reverse current status
深度优先遍历 or 递归

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    return max(left,right) + 1
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

复杂度分析：
时间复杂度:O(n) 遍历所有树的节点元素
空间复杂度:O(height) 递归时用到的栈空间,height 为二叉树的高度

广度优先遍历：类似水波纹一层层遍历：
广度优先遍历需要借助队列实现
疑惑的点：
1. 什么时间出队
2. 什么时间将元素放进队列中
3. 什么时间将当前层数加
那你能不能先实现一版 二叉树的层序遍历呢
比较卡我的地方是我如何知道这一层遍历结束了呢

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    depth := 0
    for len(queue) > 0 {
        depth ++
        size := len(queue)
        for size > 0 {
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            size--
        }
    }
    return depth
}

看完题解以后，知道了这道题自己卡在了 len(queue) 上面
自己写代码的时候弹出栈的时候将
 queue = queue[1:]
写成了
queue := queue[1:]
然后程序就变成了死循环，自己没想明白为什么会变成死循环了
后来自己打印了一下，发现 queue[0] 一直是 头节点3，也就是没有把头节点去掉
所以很小的地方，也会引起bug导致程序无法正常运行，也说明了基础知识的重要性以及对编程语言基本结构的熟练程度以及实现细节，了解的越多，出现问题的时候才可能想的清楚

复杂度分析：
时间复杂度：O(n) : 树节点的个数
空间复杂度: O(n)：队列存储元素的数量

总结：
1. 看了一下，这道题自己提交了12次了，递归是可以写出来了，用队列的时候还是有卡壳的地方
2. 学习算法和数据结构越来越觉得和学英语一样，没有捷径，需要一点一滴的积累，多用，多反复练习
