/*
给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
示例 1:
输入:       1         1
          / \       / \
         2   3     2   3
        [1,2,3],   [1,2,3]
输出: true
示例 2:
输入:      1          1
          /           \
         2             2
        [1,2],     [1,null,2]
输出: false
示例 3:
输入:       1         1
          / \       / \
         2   1     1   2
        [1,2,1],   [1,1,2]
输出: false
*/

Clearfication:
判断两个数结构相同，并且节点具有相同的值，则认为它们是相同的
什么是结构体相同呢？ 就是完全一样
如果节点具有相同的值，结构相同么？不一定的：
如
1
   2
3
和
1
    2  3
前序遍历是相同的，但是结构是不同的，所以我们需要遍历每个节点的时候都要做判断，判断节点是否相等
思路是有的：代码不知道如何组织了。。。
看了题解：

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

自己的思路卡在了 如何什么时候返回 true/false 上面了,这一步没有想清楚，看到题解里面这一行代码
return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
明白了什么时候返回值了，一个大问题划分为若干个相似的小问题，将小问题解决了，大问题也就解决了
返回值也是一样的，不知道什么时候进行返回值，让子问题去返回就好了

复杂度分析：
时间复杂度: O(n):二叉树节点的个数
空间复杂度：O(n): 空间复杂度取决于递归调用的层数，层数不会超过二叉树的高度，最坏情况下，二叉树的高度为n, 好的情况下为log(N)

那么如何使用广度优先遍历呢？
使用队列进行广度优先遍历，然后比较元素值, 参考题解：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    queue1,queue2 := []*TreeNode{p},[]*TreeNode{q}
    for len(queue1) > 0 && len(queue2) > 0 {
        node1,node2 := queue1[0],queue2[0]
        queue1,queue2 = queue1[1:],queue2[1:]
        if node1.Val != node2.Val {
            return false
        }
        left1,right1 := node1.Left,node1.Right
        left2,right2 := node2.Left,node2.Right
        if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
            return false
        }
        if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
            return false
        }
        if left1 != nil {
            queue1 = append(queue1, left1)
        }
        if left2 != nil {
            queue2 = append(queue2, left2)
        }
        if right1 != nil {
            queue1 = append(queue1, right1)
        }
        if right2 != nil {
            queue2 = append(queue2, right2) 
        }
    }
    return len(queue1) == 0 && len(queue2) == 0
}

复杂度分析：
时间复杂度：O(n)：二叉树元素个数
空间复杂度：O(n): 使用队列存储二叉树元素个数

总结：
1. 递归：大问题化解为小问题，这个大问题我当前不知道结果，然后分解为小问题，小问题的结果汇总就是大问题的结果

2. 思维要开阔，尽信书不如无书
