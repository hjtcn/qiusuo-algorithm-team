/*
给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
示例 ：
输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]
       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \ 
1        7   9
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9  
提示：
给定树中的结点数介于 1 和 100 之间。
每个结点都有一个从 0 到 1000 范围内的唯一整数值。
*/

1. Clearfication:
树中最左边的结点现在是树的根，每个结点没有左子结点，只有一个右子结点。
一开始没有理解这道题的意思，看了题解以后发现就是构造有序数组，然后根据有序数组构造一颗右下方向的斜树。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    ret := []int{}
    dfs(root,&ret)
    
    node := &TreeNode{}
    cur := node
    for _,v := range ret {
        cur.Right = &TreeNode{
            Val:v,
            Left:nil,
            Right:nil
        }
    }
    return node
}
func dfs(root *TreeNode,ret *int) {
    if root == nil {
        return
    }
    dfs(root.Left,ret)
    *ret = append(*ret, root.Val)
    dfs(root.Right,ret)
}
上面一维数组参数写错了，链表的变化代码还是不熟悉，感觉有点害怕链表的感觉
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    ret := []int{}
    dfs(root,&ret)
    
    node := &TreeNode{}
    cur := node
    for _,v := range ret {
        cur.Right = &TreeNode{
            Val:v,
            Left:nil,
            Right:nil,
        }
        cur = cur.Right
    }
    return node.Right
}
func dfs(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    dfs(root.Left,ret)
    *ret = append(*ret, root.Val)
    dfs(root.Right,ret)
}
看到方法二的时候，感觉还是蛮有意思的, 中序遍历 + 更改树的连接方式
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    ret := &TreeNode{}
    cur := ret
    inorder(root, ret)
    return ret.Right
}
func inorder(root *TreeNode,cur *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left, cur)
    root.Left = nil
    cur.Right = root
    cur = root
    inorder(root.Right,cur)
}
仿照着写出来运行结果是不对的，发现是 cur 指针我是通过函数传进去的，传进去它就会变化，变化的话效果就不一样了，所以需要将 cur 声明为全局变量
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var cur *TreeNode
func increasingBST(root *TreeNode) *TreeNode {
    ret := &TreeNode{}
    cur = ret
    inorder(root)
    return ret.Right
}
func inorder(root *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left)
    root.Left = nil
    cur.Right = root
    cur = root
    inorder(root.Right)
}
中序遍历也可以使用BFS写的

复杂度分析：
时间复杂度：O(n) 遍历每个节点的时间复杂度
空间复杂度：O(n) 递归调用栈空间

总结：
1. 越写感觉里面细节点还是很多的，如变量的定义，局部变量和参数变量的定义，这些细节点如果不注意的话，程序就会报错的，以前自己觉得，知道怎么写就行了，发现这样想挺傻的，从知道 -》实现 这个过程里面还是有很多工作要做的
