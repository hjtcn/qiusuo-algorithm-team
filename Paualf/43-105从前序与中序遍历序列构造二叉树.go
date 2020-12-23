/*
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

1. Clearfiction:

从前序遍历中我们可以找到根节点
从中序遍历中我们可以找到根节点所在的位置
在中序遍历中  第一个元素到跟节点所在的位置为左子树
根节点所在的位置到数组末尾是右子树
如根节点位置为0，根节点在中序遍历中的位置为i
则伪代码为

root := &TreeNode{}
root.Val = preorder[0]
for i := 0;i < len(inorder);i++ {
    if inorder[i] == preorder[0] {
        return 
    }
}

root.Left = buildTree(preorder[0:i], inorder[0:i])
root.Right = buildTree(preorder[i+1:],inorder[i:1:])
return root

Coding:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    root := &TreeNode{}
    root.Val = preorder[0]
    i := 0
    for ;i < len(inorder);i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[0:i], inorder[0:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root
}

写错的地方：
1. root 初始化错误
2. 没有判断边界条件
3. 构造变量遍历时下标填写出错

fix:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0],nil,nil}
    i := 0
    for ;i < len(inorder);i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1:i + 1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root
}

看到官方题解是使用len计算的下标值:
 root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
 root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

感觉这样挺容易误解，也不是很好看，里面这样计算 len(inorder[:i])
看到这种题目，第一次看到还是比较陌生和害怕的，没有思路，这应该是第二次看到，所以没有那么害怕，仔细想一下还是有思路的

迭代的思路感觉不是很常规，就没有看了，感觉是在找规律，不如递归代码方便

复杂度分析：
时间复杂度：O(n) 树节点个数
空间复杂度：O(n) 递归栈空间

总结：
1. 题目有思路，思路是对的，写的代码还是有问题。。。

2. 遇到问题，不要害怕，仔细分析一下

3. 要养成写伪代码验证逻辑
