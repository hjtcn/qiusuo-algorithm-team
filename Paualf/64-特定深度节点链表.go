给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
示例：
输入：[1,2,3,4,5,null,7,8]
        1
       /  \ 
      2    3
     / \    \ 
    4   5    7
   /
  8
输出：[[1],[2,3],[4,5,7],[8]]

1. Clearfication:
看到这道题，第一种解法是：dfs(node,level),进行前序遍历
第二种解法是 层序遍历，记录每一层节点个数，然后开始进行层序遍历
伪代码：
func dfs(node *TreeNode,level int,ret *[][]int) {
        // terminator
    if root == nil {
            return 
    }
    // process current logic
    (*ret)[level] = append((*ret)[level], node.Val)
    dfs(node.Left,level  + 1, ret)
    dfs(node.Right,level + 1, ret)
}
然后去写的时候看到函数返回值是返回的类型是 ListNode，说明自己还是没有将题目分析清楚, 自己想将返回数组的写出来
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    height := height(tree)
    ret := make([][]int, height)
    
    dfs(tree,0,&ret)
    
    fmt.Println(ret)
    return nil
}
func dfs(node *TreeNode,level int,ret *[][]int) {
        // terminator
    if node == nil {
            return 
    }
    // process current logic
    (*ret)[level] = append((*ret)[level], node.Val)
    dfs(node.Left,level  + 1, ret)
    dfs(node.Right,level + 1, ret)
}
func height(tree *TreeNode) int {
     if tree == nil {
        return 0
     }
     left := height(tree.Left)
     right := height(tree.Right)
     return max(left,right) + 1
}
func max(x,y int)int {
    if x > y {
        return x
    }
    return y
}
修改
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    list := new([]*ListNode)
    level := 0
    dfs(tree,level,list)
    return *list
}
func dfs(node *TreeNode,level int,list *[]*ListNode) {
        // terminator
    if node == nil {
            return 
    }
    if level >= len(*list) {
            *list = append(*list, &ListNode{node.Val,nil})
    }else {
        head := (*list)[level]
        for head.Next != nil {
             head = head.Next          
        }
        head.Next = &ListNode{node.Val,nil}
    }
    // process current logic
    dfs(node.Left,level  + 1, list)
    dfs(node.Right,level + 1, list)
}
然后遍历数组，将每个数组中的元素使用逗号连接起来，这个 *ListNode 返回类型自己还是有点迷惑的
问题：
1. 怎么将数组元素使用逗号,连接起来
2. 返回数组类型为 ListNode
如果是BFS数组的话怎么解决呢：
伪代码：
queue := []*TreeNode{root}
ret := [][]int{}
for len(queue) > 0 {
      size := len(queue)
    tmp := []int{}
    for i := 0;i < size;i++ {
            node := queue[0]
        queue = queue[1:]
        
        tmp = append(tmp, node.Val)
        
        if node.Left != nil {
                queue = append(queue,node.Left)
        }
        
        if node.Right != nil {
                queue = append(queue,node.Right)
        }
    }
    ret = append(ret, tmp)
}
数组输出结果是对的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    queue := []*TreeNode{tree}
    ret := [][]int{}
    for len(queue) > 0 {
        size := len(queue)
        tmp := []int{}
        for i := 0;i < size;i++ {
            node := queue[0]
            queue = queue[1:]
            
            tmp = append(tmp, node.Val)
            
            if node.Left != nil {
                queue = append(queue,node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue,node.Right)
            }
        }
        ret = append(ret, tmp)
    }
    fmt.Println(ret)
    return nil
}
修改
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    queue := []*TreeNode{tree}
    ret := make([]*ListNode,0)
    for len(queue) > 0 {
        size := len(queue)
        node := new(ListNode)
        tempNode := node
        for i := 0;i < size;i++ {
            
            
            node := queue[0]
            queue = queue[1:]
            
            tempNode.Next = &ListNode{
                    Val:node.Val,
            }
            tempNode = tempNode.Next
            
            if node.Left != nil {
                queue = append(queue,node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue,node.Right)
            }
        }
        ret = append(ret, node.Next)
    }
    
    return ret
}
1. 需要求高度就去求二叉树的高度

复杂度分析：
时间复杂度：O(n) : BFS的时间复杂度是O(n),遍历节点信息 dfs的时候要找到最后一个节点信息，所以是O(n*n),遍历节点信息及找到链表最后一个节点

空间复杂度：O(n) ： 递归调用栈空间信息，及队列信息

总结：
1. 没有分析清楚题目意思，没有看到要返回的信息

2. 这段代码还是蛮有意思的,如果是头节点，则初始化头节点，如果不是找到最后的节点，赋值给它的next
  if level >= len(*list) {
            *list = append(*list, &ListNode{node.Val,nil})
    }else {
        head := (*list)[level]
        for head.Next != nil {
             head = head.Next          
        }
        head.Next = &ListNode{node.Val,nil}
    }
