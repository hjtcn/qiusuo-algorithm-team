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

说实话，看了删除元素以后要维护二叉树的性质

我自己挺害怕的

知道自己不会写

2. 看题解：

如果 key > root.Val 说明要删除的节点在右子树，root.Right = delete(root,Right,key)

如果 key < root.Val 说明要删除的节点在左子树，root.Left = delete(root.Left,key)

如果 key == root.Val 则该节点就是我们要删除的节点，则
如果该节点是叶子节点，则直接删除它：root == nil
如果该节点不是叶子节点且有右节点，则用它的后继节点值替代 root.Val = successor.Val,然后删除后继节点
如果该节点不是叶子节点且只有左节点，则用它的前驱节点的值替代 root.Val = predecessor.Val,然后删除前驱节点
返回root


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
    // delete from the right subtree
    if key > root.Val {
        root.Right = deleteNode(root.Right,key)
    }else if key < root.Val {
    	// delete from the left subtree
        root.Left = deleteNode(root.Left,key)
    }else {
    	// the node is a leat
        if root.Left == nil && root.Right == nil {
        	root = nil
        }else if root.Right != nil {
            root.Val = successor(root)
            root.Right = deleteNode(root.Right,root.Val)
        }else {
            root.Val = predecessor(root)
            root.Left = deleteNode(root.Left, root.Val)
        }
    }
    return root
}

/*
	one step right and then always left
*/
func successor(root *TreeNode) int{
	root = root.Right
    for root.Left != nil {
    	root = root.Left
    }
    return root.Val
}

/*
	one step left and then always right
*/
func predecessor(root *TreeNode) int{
	root = root.Left
    for root.Right != nil {
    	root = root.Right
    }
    return root.Val
}

3. 复杂度分析：
时间复杂度：O(logN)
空间复杂度：O(H) H为树的高度

4. 总结
4.1: 通过这道题还是觉得自己对复杂问题害怕，不敢下手去分析
4.2: one step right and then always left  one step left and then always right，通过这两端程序，感觉数据结构真的很巧妙
4.3：感觉成长中最大的敌人其实是你自己，去了解自己，然后慢慢的改变自己，引导自己的想法和行为
4.4: 其实题目中的java题解写的蛮好的，自己还是想要去看Go的代码，其实要更加关注思路和思考过程，去看具体代码就是去关注细节关注实现了，还是太懒了，还是害怕失败，害怕自己看不懂，害怕。。。。
