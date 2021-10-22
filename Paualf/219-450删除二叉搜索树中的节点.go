给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

1. Clarification:

看到这道题第一反应是没有思路哈，不知道如何下手。。。

2. 看题解:
看完题解知道思路以后写的;连贯起来的地方没有处理好
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val > key {
        deleteNode(root.Left,key)
    }else if root.Val < key {
        deleteNode(root.Right,key)
    }else {
        if root.Right != nil {
            root = successor(root)
        }else if root.Left != nil {
            root = lefter(root)
        }else {
            root = nil
        }
    }

    return root
}

// one right then all left
func successor(root *TreeNode)*TreeNode {
    root = root.Right

    for root != nil {
        root = root.Left
    }

    return root
}

// one left all right
func lefter(root *TreeNode)*TreeNode {
    root = root.Left

    for root != nil {
        root = root.Right
    }

    return root
}

改了一下：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val > key {
        root.Left = deleteNode(root.Left,key)
    }else if root.Val < key {
        root.Right = deleteNode(root.Right,key)
    }else {
        if root.Right != nil {
            root.Val = successor(root)
            root.Right = deleteNode(root.Right,root.Val)
        }else if root.Left != nil {
            root.Val = predecessor(root)
            root.Left = deleteNode(root.Left,root.Val)
        }else {
            root = nil
        }
    }

    return root
}

// one right then all left
func successor(root *TreeNode) int {
    root = root.Right

    for root.Left != nil {
        root = root.Left
    }

    return root.Val
}

// one left all right
func predecessor(root *TreeNode) int {
    root = root.Left

    for root.Right != nil {
        root = root.Right
    }

    return root.Val
}
修改节点的地方没有搞清楚，root.Left 和 root.Right 的地方赋值变化没有处理好

然后找 predecessor 祖先的时候 和 successor 继承者的时候自己的判断条件也没有处理清楚

思路：
如果 key > root.Val ,说明要删除的节点在右子树，root.RIght = deleteNode(root.RIght,key)
如果 key < root.Val,说明要删除的节点在左子树，root.Left = delteNode(root.Left,key)
如果 key == root.Val,则该节点就是我们要删除的节点，则：
      如果该节点是叶子节点，则直接删除它：root = nil
 如果该节点不是叶子节点且有右节点，则用它的后继节点的值替代 root.val = successor.Val,然后删除后继节点
 如果该节点不是叶子节点且有左节点，则用它的前驱节点的值替代 root.Val = predecessor.Val,然后删除前驱节点
返回root

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

总结：
4.1: 自己写了一遍没有写出来，其实本质上还是自己觉得自己理解了题解，其实只是理解了60%，里面很多细节如果不自己实现的话是不知道的
