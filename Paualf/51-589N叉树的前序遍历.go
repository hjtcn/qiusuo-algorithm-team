/*
给定一个 N 叉树，返回其节点值的前序遍历。
例如，给定一个 3叉树 :
 
 
返回其前序遍历: [1,3,5,6,2,4]。
 
说明: 递归法很简单，你可以使用迭代法完成此题吗?
*/

1. Clearfiction:
前序遍历：左
二叉树的前序遍历：
*ret = append(*ret, root.Val)
dfs(root.Left)
dfs(root.Right)

如果是N叉树的话：
if root == nil {
return
}
*ret = append(*ret, root.Val)
for i := 0;i < len(root.Children);i++ {
dfs(root.Children[i])
}

DFS:
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {
    ret := []int{}
    dfs(root,&ret)
    return ret
}
func dfs(root *Node,ret *[]int) {
    if root == nil {
        return
    }
    *ret = append(*ret, root.Val)
    for i := 0;i < len(root.Children);i++ {
        dfs(root.Children[i], ret)
    }
}

一开始的时候写成了 dfs(root.Children[0],ret)，后来分析了输出结果看到了 应该是写 dfs(root.Children[i],ret) 才对
BFS 自己像二叉树的栈进行迭代的时候没有想通，没有想通的地方还是没有真正的理解，先放入结果数组，然后将节点放入到栈中
N叉树的迭代遍历，参考了题解：
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {
     res := []int{}
     stack := []*Node{root}
     for len(stack) > 0 {
         for root != nil {
             res = append(res, root.Val)
             if len(root.Children) == 0 {
                 break
             }
             for i := len(root.Children) - 1;i > 0;i-- {
                 stack = append(stack, root.Children[i])
             }
             root = root.Children[0]
         }
         root = stack[len(stack) - 1]
         stack = stack[:len(stack) - 1]
     }
     return res
}

主要精髓在
 for i := len(root.Children) - 1;i > 0;i-- {
      stack = append(stack, root.Children[i])
}

root = root.Children[0]

这几行代码，中间将栈的使用和思想运用上了，栈理解起来很简单，但是越是简单的东西使用好就越不简单
又看到这个代码，根据还是蛮有意思的，利用tmp，将最左边的元素放到了栈顶
func preorder(root *Node)[]int{
    res := []int{}
    
    if root == nil {
        return res
    }
    
    stack := []*Node{root}
    for len(stack) > 0 {
        r := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        res = append(res, r.Val)
        
        tmp := []*Node{}
        
        for _,v := range r.Children {
            tmp = append([]*Node{v}, tmp...)
        }
        stack = append(stack, tmp...)
    }
    return res
}
比较有意思的代码在
 tmp := []*Node{}
        
for _,v := range r.Children {
    tmp = append([]*Node{v}, tmp...)
}
stack = append(stack, tmp...)
还看到这样的代码,感觉更简洁，感觉是在递归
func preorder(root *Node) []int {
    if root == nil {
        return nil
    }
    ret := []int{root.Val}
    for _, child := range root.Children {
        ret = append(ret, preorder(child)...)
    }
    return ret
}

复杂度分析：
    时间复杂度：O(n) 节点数量
    空间复杂度：O(n) 递归栈空间和使用栈存储队列的空间开销

总结：
1. 对栈的使用和理解还是不是很深刻

2. 举一反三，推而广之，n 叉树的dfs是自己参考二叉树的dfs写出来的
