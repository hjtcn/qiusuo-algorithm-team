给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

例如，

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2     
     / \   
    1   3
在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。

1. Clarification:

在BST 中找到节点值等于给定值的节点。返回这个节点为根的子树。

返回子树，如果找到这个节点是不是就可以返回了？
try it

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == val {
        return root
    }
    
    left := searchBST(root.Left, val) 
    right := searchBST(root.Right, val)

    if left != nil {
        return left
    }

    if right != nil {
        return right
    }

    return nil
}

2. 看题解：

递归可以这样优化：

我上面的写法没有利用到二叉搜索树的性质
// 递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
// 迭代：for循环
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			// 根节点值大于目标值，所以选择左子树
			root = root.Left
		} else {
			// 根节点值小于等于目标值，所以选择又子树
			root = root.Right
		}
	}
	return root
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：递归：O(n), 迭代:O(1)

4. 总结：
4.1: 理解题目然后转化成自己的理解往往是很重要的一步

4.2: 不要假装勤奋，要去多思考
