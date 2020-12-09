/*
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：
输入：root = []
输出：[]
示例 3：
输入：root = [1]
输出：[1]
示例 4：
输入：root = [1,2]
输出：[2,1]
示例 5：
输入：root = [1,null,2]
输出：[1,2]
 
提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
 
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

Clarfication:
二叉树的中序遍历：
左根右  进行遍历，需要存储节点元素放入到数组中，可以传递数组地址

Possible solution:
递归  时间复杂度：O(n)遍历整个二叉树节点，空间复杂度O(height) 递归使用的栈空间
迭代算法，使用栈的数据结构进行辅助进行模拟， 时间复杂度O(n) :遍历整个二叉树节点，空间复杂度：O(n) 需要开辟栈进行辅助进行模拟

Coding
递归进行解决
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root, &ret)
    return ret
}
func helper(root *TreeNode,* ret) {
        if root == nil {
          return nil
    }
    
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}
上面的代码是不能运行的，有bug的地方在于helper 的参数 *ret
需要改成 ret *[]int, 执行数组的指针
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root, &ret)
    return ret
}
func helper(root *TreeNode,ret *[]int) {
    if root == nil {
          return
    }
    
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

复杂度分析：
时间复杂度：O(n): 遍历二叉树的每个节点
空间复杂度：O(height) : 递归时候开辟栈空间

使用栈进行迭代
思路：
找重复性：
先将根节点放入栈中，一直循环将节点的左节点放入到栈中，当左节点为空时
弹出栈中节点，放入到结果数组中，然后将遍历节点指向弹出节点的右节点，重复该过程

func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    
    if root == nil {
        return ret 
    }
 
    stack := []*TreeNode{}
 
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        
        if root == nil && len(stack) > 0 {
             node := stack[len(stack) - 1]
           stack = stack[:len(stack) - 1]
           ret = append(ret, node.Val)
           root = node.Right
        }
    }
    
    return ret
}

这次竟然蒙出来了，哈哈哈，正准备调试不出来的时候看下题解,这里加了一个循环就过了
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
循环中的这个判断是可以去掉的
if root == nil && len(stack) > 0 {
}
因为外层已经判断了 root 不为空或者 栈的长度 大于0 才开始循环，所以这个判断是不需要的

func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    
    if root == nil {
        return ret 
    }
 
    stack := []*TreeNode{}
 
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        ret = append(ret, node.Val)
        root = node.Right
    }
    
    return ret
}

复杂度分析：
时间复杂度：O(n) 遍历二叉树的每个节点
空间复杂度：O(n)  开辟存储栈空间

总结：
1. 学习算法和数据结构真的没有捷径，就是多练，多写，然后隔一段时间这道题再写一遍，和背单词是真的一个效果的，这道题一开始没有思路，看题解，后面自己有思路了，自己写，写的时候遇到问题，再看题解，如此循环，才能学会，到最后运用到工程和项目中，也是需要自己多思考和总结的

2. 不要害怕麻烦
