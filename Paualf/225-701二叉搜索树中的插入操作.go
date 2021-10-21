给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

示例 1：
输入：root = [4,2,7,1,3], val = 5
输出：[4,2,7,1,3,5]
解释：另一个满足题目要求可以通过的树是：

示例 2：

输入：root = [40,20,60,10,30,50,70], val = 25
输出：[40,20,60,10,30,50,70,null,null,25]
示例 3：

输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
输出：[4,2,7,1,3,5]
 
提示：

给定的树上的节点数介于 0 和 10^4 之间
每个节点都有一个唯一整数值，取值范围从 0 到 10^8
-10^8 <= val <= 10^8
新值和原始二叉搜索树中的任意节点值都不同

1.Clarification:
Q: 怎么将值插入二叉搜索树中？
A: 

Q：插入后会有什么变化？怎么处理这些变化
A:

如果将要插入的节点 > root.Val,则向右边插入
insertInfoBST(root.Right,val)

如果将要插入的节点< root.Val,则向左边插入
insertInfoBST(root.Left,val)

// 这一块没有想好
如果 root == nil 则将元素新建，然后写入数据

2. 看题解：

看完题解发现，后面自己的思路是正确的，所以还是尽可能的勇敢去尝试

递归：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val:val}
    }

    if val > root.Val {
       root.Right =  insertIntoBST(root.Right,val)
    }

    if val < root.Val {
        root.Left =  insertIntoBST(root.Left, val)
    }

    return root
}
一开始的时候自己写的直接 return 掉了

if val > root.Val {
return insertIntoBST(root.RIght,val)
}

if  val < root.Val {
return insertInfoBST(root.Left,val)
}

迭代
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val:val}
    }

    p := root

    for p != nil {
        if p.Val > val {
            if p.Left == nil {
                p.Left = &TreeNode{Val:val}
                break
            }
            p = p.Left
        }

        if p.Val < val {
            if p.Right == nil {
                p.Right = &TreeNode{Val:val}
                break
            }
            p = p.Right
        }
    }

    return root
}

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4.总结
4.1: 仔细想一下，如果有思路了，就动手哈，不要害怕写不出来，编程本身就是一个实践和理论结合的东西

