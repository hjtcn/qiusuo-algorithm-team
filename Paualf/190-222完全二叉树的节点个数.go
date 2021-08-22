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

完全二叉树的遍历：

我想到了二叉树的遍历，前序，中序，后序，每个都可以使用递归或者迭代去遍历二叉树

还想了下这道题既然是考察完全二叉树，那么完全二叉树的性质是什么呢？
除了最底层结点可能没被填满外，其余每层结点数都达到最大值

题目里面还给了进阶说明：
遍历树统计节点是一种时间复杂度为O(n) 的简单解决方案，可以设计一个更快的算法吗？

感觉这里应该是有一些特殊的解法的，我目前没有想到哈


还想了下，二叉树里面的前序、中序、后序 迭代的时候自己比较容易出错还是需要多注意下的

自己写了后序遍历的迭代解法  和  层序遍历的解法
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    return helper(root)
}

func helper(root *TreeNode) (ret int){
    if root == nil {
        return 0
    }

    left := helper(root.Left)
    right := helper(root.Right)

    ret = left + right + 1
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
        node := queueNode[0]
        ret++

        queueNode = queueNode[1:]

        if node.Left != nil {
            queueNode = append(queueNode, node.Left)
        }

        if node.Right != nil {
            queueNode = append(queueNode, node.Right)
        }
    }

    return ret 
}


2. 看题解：
说实话，题解里面的思路感觉还是很nice的

对完全二叉树的最后一层节点进行二分查找

如果完全二叉树的层高为 h
则最后一层节点为1个的时候 2 ^ h-1 + 1 = 2 ^h
当最后一层节点满的时候 ： 2 ^h + 1-1

所以最后一层节点个数为 [ 2 ^ h - 2^h + 1 - 1]

然后判断节点k是否存在，将k转化为二进制位，有点 哈夫曼编码的意思

思路是理解了，代码没有看懂。。。

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
        node := root
        for node != nil && bits > 0 {
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


后来搜了下 sort.Search 表示：
func Search(n int,f func(int) bool) int

采用二分法搜索找到 [0,n）区间内最小的满足 f(i) == true 的值 i。

里面的二进制还是有点难搞，哈哈哈。。。

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：递归栈 O(n), 存储节点变化也是 O(n)

总结：
4.1: 利用数据结构的性质去解决问题，题解中使用二分查找在完全二叉树最后一层查找思路还是很nice的
