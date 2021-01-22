/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
 
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/

1. Clearfication:       
     路径终止条件
dfs 思路：
if root.Left == nil && root.Right == nil {
    if left == 0 {
        tmp := make([]int, len(tmp))
        ret = append(ret, tmp)
    }
}
dfs:
func pathSum(root *TreeNode, sum int) [][]int {
     paths := [][]int{}
     path := []int{}
    
    helper(root,sum,&paths,path)
    return paths
}
func helper(root *TreeNode,sum int,paths *[][]int,path []int) {
    if root == nil {
        return
    }
    
    // terminator
    if root.Left == nil && root.Right == nil && sum == 0 {
        tmp := make([]int, len(path))
        copy(path,tmp)
        *paths = append(*paths,tmp)
        return
    }
    
    // process current logic
    path = append(path, root.Val)
    
    helper(root,sum - root.Val,paths,path)
    helper(root,sum - root.Val,paths,path)
}
BFS思路：
一层一层遍历，终止条件也是
root.Left == nil && root.Right == nil && sum == sum {
    
}
这道题如果用BFS的话需要维护两个队列，一个是节点变化的队列，一个是值变化的队列
值变化的队列我有点没想明白的是，我弹出一个节点，弹出几个值呢，有的时候需要弹出一个值，有的时候需要弹出两个，怎么出来弹出值的个数呢
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
        
}
func helper(root *TreeNode,sum,currentSum int,paths *[][]int,path []int) {
    currentSum := 0
    
}
上面就是我的思考思路，dfs，提交是有报错的，BFS的思路没有想清楚
DFS修改
func pathSum(root *TreeNode, sum int) [][]int {
     paths := [][]int{}
     path := []int{}
    
    helper(root,sum,&paths,path)
    return paths
}
func helper(root *TreeNode,sum int,paths *[][]int,path []int) {
    if root == nil {
        return
    }
    
    // terminator
    if root.Left == nil && root.Right == nil && sum == 0 {
        tmp := make([]int, len(path))
        copy(path,tmp)
        *paths = append(*paths,tmp)
        return
    }
    
    // process current logic
    path = append(path, root.Val)
    
    helper(root,sum - root.Val,paths,append(append([]int{}, path...)))
    helper(root,sum - root.Val,paths,append(append([]int{}, path...)))
}
改了下还是错了
func pathSum(root *TreeNode, sum int) [][]int {
     paths := [][]int{}
     path := []int{}
    
    helper(root,sum,&paths,path)
    return paths
}
func helper(root *TreeNode,sum int,paths *[][]int,path []int) {
    if root == nil {
        return
    }
    
    // terminator
    if root.Left == nil && root.Right == nil && sum == root.Val {
        path = append(path, root.Val)
        tmp := make([]int, len(path))
        copy(tmp,path)
        *paths = append(*paths,tmp)
        return
    }
    // process current logic
    if root.Left != nil {
         helper(root.Left,sum - root.Val,paths,append(append([]int{}, path...),root.Val))
    }
    if root.Right != nil {
         helper(root.Right,sum - root.Val,paths,append(append([]int{}, path...),root.Val))
    }
}
func pathSum(root *TreeNode, sum int) [][]int {
     paths := [][]int{}
     path := []int{}
    
    helper(root,sum,&paths,path)
    return paths
}
func helper(root *TreeNode,sum int,paths *[][]int,path []int) {
    if root == nil {
        return
    }
    
    // terminator
    if root.Left == nil && root.Right == nil && sum == root.Val {
        path = append(path, root.Val)
        tmp := make([]int, len(path))
        copy(tmp,path)
        *paths = append(*paths,tmp)
        return
    }
    // process current logic
    path = append(path,root.Val)
    if root.Left != nil {
         helper(root.Left,sum - root.Val,paths,path)
    }
    if root.Right != nil {
         helper(root.Right,sum - root.Val,paths,path)
    }
}
这样也是可以的
上面是对的，我们来复盘下自己那里写错了
判断终止条件的时候做错了，我写的判断条件是 sum == 0, 如果 节点的左节点为空并且节点的右节点为空，那么我们就处理这个节点并判断这个节点是否满足题目条件，如果走下去的话是有报错的，我们有判断root 是否等于nil的逻辑，如果成立，函数直接return 退出，这里还是蛮细节的，个人感觉比较容易出错的地方
BFS呢，没有想通，维护两个队列的时候，第二个队列的值怎么变化的
BFS 自己疑惑的地方，其实还是没那么复杂的，从队列节点中弹出一个，然后从值队列中弹出一个，值队列中的值 + 队列节点的left 或者 队列节点的 right
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    queNode := []*TreeNode{}
    queVal := []int{}
    queNode = append(queNode, root)
    queVal = append(queVal, root.Val)
    
    ret := [][]int{}
    path := []{}
    for len(queNode) != 0 {
        now := queNode[0]
        queNode = queNode[1:]
        temp := queVal[0]
        queVal = queVal[1:]
        if now.Left == nil && now.Right == nil {
            if temp == sum {
                path = append(path, temp)
                tmp := make([]int, len(path))
                copy(tmp, path)
                ret = append(ret, tmp)
                return
            }
        }
        
        path = append(path, temp)
                      
        if now.Left != nil {
            queNode = append(queNode, now.Left)
            queVal = append(queVal, now.Left.Val + temp)
        }
        
        if now.Right != nil {
            queNode = append(queNode, now.Right)
            queVal = append(queVal, now.Right.Val + temp)
        }
    }
    return false
}
我用两个队列然后加了一个path 试了一下，感觉怪怪的，path的值是一直往里增加的，所以感觉这个方法不太行
题解里面使用BFS的时候，定义了这样的一个数据结构
type pair struct {
    node *TreeNode
    left int
}
然后BFS遍历的时候，记录了节点的父节点信息，等到如果符合结果，然后将结果数组找到然后返回
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type pair struct {
    node *TreeNode
    left int
}
func pathSum(root *TreeNode,sum int)[][]int {
    if root == nil {
        return nil
    }
    
    parent := map[*TreeNode]*TreeNode{}
    paths := [][]int{}
    
    queue := []pair{{root, sum}}
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        node := p.node
        left := p.left - node.Val
        
        if node.Left == nil && node.Right == nil {
            if left == 0 {
                paths = append(paths, getPath(node,parent))
            }
        }else {
            if node.Left != nil {
                parent[node.Left] = node
                queue = append(queue, pair{node.Left, left})
            }
            
            if node.Right != nil {
                parent[node.Right] = node
                queue = append(queue, pair{node.Right,left})
            }
        }
    }
    return paths
}
func getPath(node *TreeNode,parent map[*TreeNode]*TreeNode) []int{
    path := []int{}
    
    for ;node != nil;node = parent[node] {
        path = append(path,node.Val)
    }
    
    for i,j := 0,len(path)-1;i<j;i++ {
        path[i],path[j] = path[j],path[i]
        j--
    }
    return path
}
自己构建一个数据结构，去使用queue,然后使用parent map存储节点信息 找到结果了时候，然后找到数组信息返回,还是蛮巧妙和厉害的

复杂度分析：
时间复杂度：O(n): O(n) 遍历每个节点

空间复杂度：O(n): 使用dfs，递归调用栈空间，使用BFS,开辟队列信息以及map parent信息

总结：
1. 遇到问题能不能自己去构建数组结构去解决？ 培养自己分析问题，解决问题，然后使用数据结构和算法解决问题的能力
