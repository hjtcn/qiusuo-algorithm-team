给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

 

示例 1：


输入：root = [1,2,3,4,5,6]
输出：6
示例 2：

输入：root = []
输出：0
示例 3：

输入：root = [1]
输出：1
 

提示：

树中节点的数目范围是[0, 5 * 104]
0 <= Node.val <= 5 * 104
题目数据保证输入的树是 完全二叉树


进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

1. Clarification:

这道题目还是考察完全二叉树的性质的，从左一直查找可以获取到它的层数

Q:完全二叉树的叶子节点是靠树的右边，但是怎么确定最后一层树右边节点所在的位置呢？
A：这个没想清楚哈

二叉树的节点个数，后序遍历
coding:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    ret := 0
    helper(root,&ret)
    return ret
}

func helper(root *TreeNode,ret *int) {
    if root == nil {
    	return
    }
    helper(root.Left,ret)
    helper(root.Right,ret)
    *ret = *ret + 1
    return
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queueNode := []*TreeNode{root}
    ret := 0
    
    for len(queueNode) > 0 {
        size := len(queueNode)
        
            for size > 0 {
                size--
                node := queueNode[0]
                queueNode = queueNode[1:]
          		ret++
                
                if node.Left != nil {
                    queueNode = append(queueNode,node.Left)
                }
                
                if node.Right != nil {
                    queueNode = append(queueNode,node.Right)
                }
            }
    }
    
    return ret
}

2. 看题解：
题解还是蛮巧妙的，在最后一层使用二分查找寻找最右边的节点
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    level := 0
    for node := root; node.Left != nil; node = node.Left {
        level++
    }
    return sort.Search(1<<(level+1), func(k int) bool {
        if k <= 1<<level {
            return false
        }
        bits := 1 << (level - 1)
       // fmt.Println(level,bits)
        node := root
        for node != nil && bits > 0 {
            // 这里自己打了断点查看了二进制的运行
            fmt.Println(bits,k,bits&k)
            if bits&k == 0 {
                node = node.Left
            } else {
                node = node.Right
            }
            bits >>= 1
        }
        return node == nil
    }) - 1
}


sort.Search 其实本质上是在用左右节点搜索空间找到第一个存在的值

一开始看go的题解没看懂，看java的理解了

3. 复杂度分析：
时间复杂度：O(n) or O(logN)
空间复杂度：O(n)

总结：
4.1: 成长真的没有啥捷径，做一道题，看一本书，做一件事情，就这样，你需要付出自己的耐心，精力，时间，唯此而已
