翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

1.Clarification:

思考了一下：

这道题目考察的是二叉树的镜像对称：怎么分解子问题呢？自己没有想出来。。。

2. 看题解：

从根节点开始，递归地对树进行遍历，并从叶子节点先开始翻转。
如果当前遍历到的节点root的左右两棵子树都已经翻转，那么我们只需要交换两棵子树的位置，即可完成以root为根节点的整棵子树的翻转。

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
    	return nil
    }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

迭代：
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		cur.Left, cur.Right = cur.Right, cur.Left

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return root
}

3. 复杂度分析:
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 自己把问题还是想的太大了，其实把问题拆分的小一点，更容易思考，这道题其实就是翻转每个节点的左右节点，然后就可以了，然后自己还是没有get到核心，没有分析出问题的关键点




