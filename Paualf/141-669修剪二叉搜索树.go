给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
示例 1：
输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
示例 2：
输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]
示例 3：
输入：root = [1], low = 1, high = 2
输出：[1]
示例 4：
输入：root = [1,null,2], low = 1, high = 3
输出：[1,null,2]
示例 5：
输入：root = [1,null,2], low = 2, high = 4
输出：[2]
 
提示：
树中节点数在范围 [1, 104] 内
0 <= Node.val <= 104
树中每个节点的值都是唯一的
题目数据保证输入是一棵有效的二叉搜索树
0 <= low <= high <= 104

1. Clearfication:
给了 low,和high 还是二叉搜索树，应该是要利用二叉搜索树的性质的
怎么利用二叉搜索树的性质进行裁剪呢？
没什么思路。。。

2. Coding:

3. 看题解：
func trimBST(root *TreeNode,L int,R int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if root.Val < L {
        return trimBST(root.Right,L,R)
    }
    
    if root.Val > R {
        return trimBST(root.Left,L,R)
    }
    
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right,L,R)
    return root
}   

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val < L {
        return trimBST(root.Right, L, R)
    }
    if root.Val > R {
        return trimBST(root.Left, L, R)
    }
    return &TreeNode{
        Val:   root.Val,
        Left:  trimBST(root.Left, L, R),
        Right: trimBST(root.Right, L, R),
    }
}

4. 复杂度分析：
时间复杂度：O(n) : 节点个数
空间复杂度：O(n):  递归栈深度

5. 总结：
5.1: 感觉挺像后序遍历的
5.2: 做了很多树的题，以为自己会了，发现还是不会。。。
5.3: 感觉如果写的代码很复杂，80%是自己没有想清楚，工作中也是这样，如果想清楚的话，代码可能不需要那么复杂的
