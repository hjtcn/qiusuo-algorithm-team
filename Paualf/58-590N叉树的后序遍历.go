/*
给定一个 N 叉树，返回其节点值的后序遍历。
例如，给定一个 3叉树 :
返回其后序遍历: [5,6,3,2,4,1].
说明: 递归法很简单，你可以使用迭代法完成此题吗?
*/

1. Clearfication:
后序遍历：
左
右
处理结果
N叉树和二叉树不一样的地方在于处理子节点的时候，二叉树里面有 左节点和右节点
N叉树里面是 Children 节点是不一样的地方，处理Children 是我们需要考虑的东西
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func postorder(root *Node) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}
func helper(root *Node,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left)
    helper(root.Right)
    ret = append(ret, root.Val)
}

写二叉树的后序遍历的时候还是报错了
func postorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}
func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    helper(root.Right,ret)
    ret = append(ret, root.Val)
}

Line 21: Char 17: first argument to append must be slice; have *[]int (solution.go)
一开始是意味自己数组指针参数又写错了，仔细看了一下是指针使用错了
func postorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}
func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    helper(root.Right,ret)
    *ret = append(*ret, root.Val)
}

忽然有点蒙了，不知道怎么写了
就看了题解：
func postorder(root *Node) []int {
    ret := []int{}
    dfs(root,&ret)
    return ret
}
func dfs(root *Node,ret *[]int) {
    if root == nil {
        return
    }
    
    for _,n := range root.Children {
        dfs(n,ret)
    }
    *ret = append(*ret, root.Val)
}

dfs迭代，使用栈迭代以后，然后交换元素,感觉如果我自己遇到这种问题，不一定可以想的出来
func postorder(root *Node) []int {
    ret := []int{}
    if root == nil {
        return ret
    }
    
    stack := []*Node{root}
    
    for len(stack) > 0 {
        root = stack[len(stack) - 1]
        ret = append(ret, root.Val)
        stack = stack[:len(stack) - 1]
        l := len(root.Children)
        for i := 0;i < l;i++ {
            stack = append(stack, root.Children[i])
        }   
    }
    
    l := len(ret) - 1
    for i := 0;i < l / 2 + 1;i++ {
        ret[i],ret[l-i] = ret[l-i],ret[i]
    }
    return ret
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间负载度：O(n)  使用队列存储元素或者递归调用栈空间

总结：
1. 自己对这种复杂度多的东西，如N叉树还是不太熟悉，可能对二叉树的遍历也不太熟悉

2. 如果特别忙的话，脑子确实可能转不动的。。。
