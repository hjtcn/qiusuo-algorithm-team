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

思路：
我们可以从后序遍历最后一位元素，定位到根节点
然后在中序遍历中找到根节点的位置，根节点左边是左子树，右边是右子树
然后重复上面这个过程，构建树结构

细节点在于后序遍历里面定位到和中序序列对应的元素

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

    pos := findInorderPos(inorder,postorder)
    //fmt.Println(pos)
    
    lInorder := inorder[0:pos]
    rInorder := inorder[pos + 1:]

    lPostOrder := postorder[0:len(lInorder)]
    rPostOrder := postorder[len(lInorder):len(lInorder) + len(rInorder)]

    //fmt.Println(len(lInorder),len(rInorder),lInorder,rInorder,lPostOrder,rPostOrder)

    root := &TreeNode{
        Val:inorder[pos],
    }

    root.Left = buildTree(lInorder,lPostOrder)
    root.Right = buildTree(rInorder,rPostOrder)

    return root
}

func findInorderPos(inorder []int,postorder []int)(pos int) {
    last := postorder[len(postorder) - 1]

    for i := 0;i < len(inorder);i++ {
        if inorder[i] == last {
            pos = i
        }
    }

    return pos
}

2. 看题解：

使用map 去记录

说实话有时候看文字题解的时候真的是在考验自己

考验自己的耐心


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n(

4. 总结：
4.1: 找bug，定位问题，不是一件容易的事情，成长本身就不是一件容易的事情哈
4.2: 培养耐心 keep patient
