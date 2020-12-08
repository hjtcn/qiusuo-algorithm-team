/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。
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

Clearfiction:
理解题目意思：从根节点到叶子节点，路径总和等于给定目标和的路径
想到的是利用递归解决，遍历到叶子节点，什么时候是叶子节点呢
if root.Left == nil || root.Right == nil {
    return 
}
递归模版：
func recursion(lenvl,n int) {
    // terminator
    if level > n {
        return
    }
    
    // process current logic
    
    // drill down
}
对比当前题目的递归模版：
路径信息需要携带
terminator: 
    root == nil
if root == nil && currentSum == sum {
    
}
// process current logic:
currentSum += root.Val
tmp = append(tmp)
// drill down
helper(root.Left,currentSum,sum,tmp)
helper(root.Right,curretSum,sum,tmp)
helper(root *TreeNode,currentSum,sum,tmp[] int) {
        
}

Coding:
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
    currentSum := 0
    tmp := []int{}
    helper(root,currentSum,sum,tmp,&ret)
    
    return ret
}

func helper(root *TreeNode,currentSum int,sum int,tmp []int, ret *[][]int) {
    // terminator
    if root == nil && currentSum == sum {
        *ret = append(*ret, tmp)
        return
    }
    // process current logic
    currentSum += root.Val
    tmp = append(tmp, root.Val)
    helper(root.Left,currentSum,sum,tmp,ret)
    helper(root.Right,currentSum,sum,tmp,ret)
}

然后运行报错了：

思路是清晰的，但是转化为代码上面，自己写的话有几个地方还是容易出错的，也就是自己代码 Ac 不了的原因
1. 如果存储二维数组的结果
2. 如何记录每层遍历一维数组的记录，也就是经过的路径节点
3. 终止条件

以上3个就是有可能写不出来的点
然后就是看题解了：看题解的时候不要走马观花，遇到好的去理解，去写，不要遇到不会的就要逃跑

func pathSum(root *TreeNode,sum int)[][]int {
    ret := [][]int{}
    path := []int{}
    dfs(&ret,root,path,sum)
    return ret
}

func dfs(ret *[][]int,root *TreeNode,path []int, target int) {
    switch {
        case root == nil:
            return
        case root.Left == nil && root.Right == nil && root.Val == target:
            dst := make([]int, len(path) + 1)
            copy(dst, append(path, root.Val))
            *ret = append(*ret, dst)
            return
    }
    path = append(path, root.Val)
    dfs(ret,root.Left,path,target - root.Val)
    dfs(ret,root.Right,path,target - root.Val)
}

换成if 的方式
func pathSum(root *TreeNode,sum int)[][]int {
    ret := [][]int{}
    path := []int{}
    dfs(&ret,root,path,sum)
    return ret
}

func dfs(ret *[][]int,root *TreeNode,path []int, target int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && root.Val == target {
            dst := make([]int, len(path) + 1)
            copy(dst, append(path, root.Val))
            *ret = append(*ret, dst)
            return
    }
    path = append(path, root.Val)
    dfs(ret,root.Left,path,target - root.Val)
    dfs(ret,root.Right,path,target - root.Val)
}

如果换成自己的思路 currentSum 和 sum 的方式
func pathSum(root *TreeNode,sum int)[][]int {
    ret := [][]int{}
    path := []int{}
    currentSum := 0
    dfs(&ret,root,path,currentSum, sum)
    return ret
}

func dfs(ret *[][]int,root *TreeNode,path []int, currentSum,sum int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && root.Val + currentSum == sum {
            dst := make([]int, len(path) + 1)
            copy(dst, append(path, root.Val))
            *ret = append(*ret, dst)
            return
    }
    path = append(path, root.Val)
    dfs(ret,root.Left,path,root.Val + currentSum, sum)
    dfs(ret,root.Right,path,root.Val + currentSum, sum)
}

题解里面还有BFS:自己思考的时候也想了BFS,没有想到思路：

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
        for ; node != nil;node = parent[node] {
            path = append(path, node.Val)
        }
        
        for i,j := 0,len(path) - 1;i < j;i++ {
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
                ans = append(ans,getPath(node))
            }
        }else {
            if node.Left != nil {
                parent[node.Left] = node
                queue = append(queue,pair{node.Left, left})
            }
                               
             if node.Right != nil {
                parent[node.Right] = node
                queue = append(queue, pair{node.Right,left})
             }
        }
    }
                               
    return
}

BFS感觉还是有意思的，和以前BFS不同的是，queue 里面存储的是 结点和当前需要匹配的值，然后存储 parent 的路径，将parent 的路径放入到最后结果数组里面

复杂度分析：
  时间复杂度：题解里面写的是O(N*N),我个人觉得是O(N),因为是dfs，比那里每个节点，时间复杂度为O(N)
  空间复杂度：O(N) 递归调用栈空间大小，然后BFS的时候建立队列存储节点关系

总结：
1. 慢慢开始从DFS演变为递归了，难度慢慢有了，自己要多花一些时间了

2. 不要害怕，将题目分析清楚，将知识点拆碎揉烂学会它
