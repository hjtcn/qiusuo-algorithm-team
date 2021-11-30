根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

1. Clarification:
1.1: 从后序遍历中可以得到根节点
1.2: 然后在中序遍历中根据根节点的元素，在中序遍历中分开，左边是左子树，右边是右子树
1.3: 根据第2步的元素个数分离后序遍历的元素，继续这个循环构造二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }

    // 获取根节点元素
    rootVal := postorder[len(postorder) - 1]
    root := &TreeNode{
        Val:rootVal,
    }

    inorderPos := findInorder(inorder,rootVal)
    root.Left = buildTree(inorder[0:inorderPos],postorder[0:inorderPos])
    root.Right = buildTree(inorder[inorderPos+1:],postorder[inorderPos:len(postorder)-1])

    return root
}

// 在中序遍历的数组中找根节点所在的位置
func findInorder(inorder []int,val int)(pos int) {
    for k,v := range inorder {
        if v == val {
            return k
        }
    } 
    return
}
2. 看题解：
可以将查找元素的过程放到hash表中，进行快速查找


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

4. 总结：
4.1: 方案写清楚了，代码才能写清楚哈
