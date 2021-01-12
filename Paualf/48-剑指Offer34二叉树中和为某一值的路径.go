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
a. 记录路径及路径经过节点的和
b. 终止条件，叶子节点，并且和为sum
c. 返回二维数组
伪代码：
判断叶子节点:

if root.Left == nil && root.Right == nil {
    
}

终止条件：
if root.Left == nil && root.Right == nil {
    if currentSum == sum {
        tmp := make([]int,len(path))
        copy(tmp, path)
        
        *ret = append(*ret, path)
        return
    }
}

DFS:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    ret := [][]int{}
    path := []int{}
    currentSum := 0
    helper(root,&ret,path,currentSum,sum)
    return ret
}
func helper(root *TreeNode, ret **[]int,path []int,currentSum,sum int) {
    // terminator
    if currentSum == sum {
        tmp := make([]int, len(path))
        copy(tmp, path)
        
        *ret = append(*ret, path)
        return
    }
    
    // process current logic & drill down
    helper(root.Left,ret,append(path, root.Val),currentSum + root.Val, sum)
    
    helper(root.Right,ret,append(path,root.Val),currentSum + root.Val, sum)
}

还是报错了
Line 5: Char 11: cannot use ret (type [][]int) as type **[]int in argument to helper (solution.go)
Line 15: Char 22: first argument to append must be slice; have *[]int (solution.go)
哈哈哈，先不看题解，看看BFS的代码可以通过么

BFS的代码要注意的地方
1. 使用队列存储结点的变化
2. 使用队列存储路径的变化

自己不会的地方：
1. 发现BFS的代码里面路径变化这部分的代码还是没有想清楚是怎么变化的
2. 然后不知道怎么将路径放入到结果数组中
BFS 不会写

然后DFS报错的地方：
1. 指针还是写错了
ret **[]int 应该是  ret *[][]int // 指向二维数组的指针

逻辑也是错的
func pathSum(root *TreeNode, sum int) [][]int {
    ret := [][]int{}
    path := []int{}
    currentSum := 0
    helper(root,&ret,path,currentSum,sum)
    return ret
}
func helper(root *TreeNode,ret *[][]int,path []int,currentSum,sum int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && root.Val + currentSum == sum {
        tmp := make([]int, len(path)+1)
        copy(tmp, append(path, root.Val))
        *ret = append(*ret, tmp)
        return
    }
    path = append(path, root.Val)
    helper(root.Left,ret,path,root.Val + currentSum, sum)
    helper(root.Right,ret,path,root.Val + currentSum, sum)
}

BFS:
使用数组记录路径信息
type pair struct {
    node *TreeNode
    left int
}
func pathSum(root *TreeNode,sum int)(ans [][]int) {
    if root == nil {
        return
    }
    
    parent := map[*TreeNode]*TreeNode{}
    
    getPath := func(node *TreeNode)(path []int) {
        for ;node != nil;node = parent[node] {
            path = append(path, node.Val)
        }
        
        for i,j := 0,len(path)-1;i < j;i++ {
            path[i],path[j] = path[j],path[i]
            j--
        }
        return
    }
    
    queue := []pair{{root, sum}}
    
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        node := p.node
        left := p.left - node.Val
        
        if node.Left == nil && node.Right == nil {
            if left == 0 {
                ans = append(ans, getPath(node))
            }
        }else {
            if node.Left != nil {
                parent[node.Left] = node
                queue = append(queue,pair{node.Left, left})
            }
            
            if node.Right != nil {
                parent[node.Right] = node
                queue = append(queue, pair{node.Right, left})
            }
        }
    }
    
    return
}

BFS:
使用parent 数组存储节点路径
使用pair 节点和剩余元素 组成队列元素

复杂度分析：
时间复杂度：O(n): 遍历每个节点 
空间复杂度：O(n)：O(n) 递归栈空间和迭代存储队列空间

总结：
1. 总结、思考、抽象、分析、Coding

2. 做过的题目可能还是不会做，第一次做要注意整理结构，细节点可能注意不到，再遇见的类似的情况不仅要注意整体结构同时要注意细节点
