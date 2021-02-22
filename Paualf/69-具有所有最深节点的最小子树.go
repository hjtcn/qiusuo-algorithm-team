给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。
如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
一个节点的 子树 是该节点加上它的所有后代的集合。
返回能满足 以该节点为根的子树中包含所有最深的节点 这一条件的具有最大深度的节点。
注意：本题与力扣 1123 重复：https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/
示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：
我们返回值为 2 的节点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的节点。
注意，节点 5、3 和 2 包含树中最深的节点，但节点 2 的子树最小，因此我们返回它。
示例 2：
输入：root = [1]
输出：[1]
解释：根节点是树中最深的节点。
示例 3：
输入：root = [0,1,3,null,2]
输出：[2]
解释：树中最深的节点为 2 ，有效子树为节点 2、1 和 0 的子树，但节点 2 的子树最小。
 
提示：
树中节点的数量介于 1 和 500 之间。
0 <= Node.val <= 500
每个节点的值都是独一无二的。

1. Clearfication:
分解问题：
a. 找到节点的深度最大的节点 =》求深度 也就是高度
b. 以该节点为根的子树中包含所有最深的节点 =》 包含所有最深的节点
题目中考察的东西还是蛮多的：       
子树的概念： 一个节点的子树是该节点加上它的所有后代的集合
以该节点为根的子树中包含所有最深的节点
1. 找到最大的深度的结点
2. 返回其父节点的值及其兄弟节点
哈哈哈，还是有点复杂的，自己没有分析清楚，就去看题解了
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    h1,h2 := 0,0
    if root.Left != nil {
        h1 = getDepth(root.Left)
    }
    
    if root.Right != nil {
        h2 = getDepth(root.Right)
    }
    
    if h1 > h2 {
        return subtreeWithAllDeepest(root.Left)
    }else if h1 < h2 {
        return subtreeWithAllDeepest(root.Right)
    }else {
        return root
    }
}

func getDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }else {
        return max(getDepth(root.Left),getDepth(root.Right)) + 1
    }
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    _, ret := subtree(root)
    return ret
}
func subtree(root *TreeNode) (int, *TreeNode) {
    if root == nil {
        return 0, nil
    }
    leftDepth, left := subtree(root.Left) // 左深度
    rightDepth, right := subtree(root.Right) // 右深度
    if leftDepth == rightDepth { // 如果左右相等，说明符合要求，深度加一
        return leftDepth + 1, root
    }
    // 如果不想等，只能返回左右节点中较深的一个
    if leftDepth > rightDepth { 
        return leftDepth + 1, left
    }
    return rightDepth + 1, right
}

如果左右子树的深度，就返回的是当前父节点，否则谁的深度更深就把谁给返回去

上面都是题解中的代码，分析一下时间复杂度

复杂度分析：

时间复杂度：
	第一种方法应该是O(n*n),每个节点遍历一次时间复杂度是O(n),然后计算一下getDepth的时间复杂度也是O(n),所以是O(n*n),也就是自上而下的遍历
        第二种方法是 后序遍历，每个节点遍历一次，所以时间复杂度是O(n)

空间复杂度：
	第一种方法自上而下的空间复杂度是O(nlogn),最坏情况下，树是一个斜树，这个时候空间复杂度是O(n*n)
	第二种方法是自下而上的遍历，也就是尾递归，它的空间复杂度是O(n)
总结：
1. 题目还是比较难理解的，感觉和日常题解需求一样，产品说的，可能并不是他想要的，你要多问一下为什么要这样做？可能才知道做这件事情的前因后果

2. 简单的二叉树遍历，自上而下和自下而上的我们都熟悉了，换到其他题目中还是没有那么熟悉，说明还是对二叉树的遍历没有真正的熟悉和掌握
