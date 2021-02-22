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
1. Clearfication:
 分析题意：
方法1：递归对二叉树进行深度优先遍历，如果最后结果为 sum = 22,将路径加入到二维结果数组当中
方法2: 使用BFS 对二叉树进行广度优先遍历，queue 里面存储的是当前sum 和 设计存储结果记录它走过的路径

伪代码：
递归模版
func recursion(level,max int,params []int) {
    // terminator
    if level > max {
        return
    }
    // process current
    
    
    // drille down
    recursion(level +1,max,new params)
}

paths := [][]int{}
func dfs(root *TreeNode,paths [][]int,path []int, target) {
    // terminator
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && target == root.Val {
        tmp := make([]int, len(path))
        copy(tmp, path)
        paths = append(paths, tmp)
        return
    }
    // process current logic
    dfs(root,paths,path,target - root.Val)
    dfs(root,paths,path,target - root.Val)
    // drill down
}

BFS:
队列存储设计：
type Q {
    sum int,
    root *TreeNode,
}
如何记录存储路径：
记录下来存储路径以后，最后如何根据结果将路径找出并打印出来
上面这些细节点才是自己真正不会的原因，也是不敢去写的原因吧，So 还是写的代码太少，代码能力不高

2. Coding
DFS:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    paths := [][]int{}
    path := []int{}
    
    dfs(root,paths,path,targetSum)
    
    return paths
}

func dfs(root *TreeNode,paths [][]int,path []int,targetSum int) {
    // terminator
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && root.Val == targetSum {
        path = append(path, root.Val)
        tmp := make([]int, len(path))
        copy(tmp, path)
        paths = append(paths, tmp)
        return
    }
    
    // process current logic
    path = append(path, root.Val)
    dfs(root,paths,path,targetSum - root.Val)
    dfs(root,paths,path,targetSum - root.Val)
}

没想清楚的点： paths 和 path 在dfs 参数里面如何传递呢？
paths 还是应该传地址的，函数里面的参数自己又忘了，指向二维数组的指针
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    paths := [][]int{}
    path := []int{}
    
    dfs(root,&paths,path,targetSum)
    
    return paths
}

func dfs(root *TreeNode,paths *[][]int,path []int,targetSum int) {
    // terminator
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && root.Val == targetSum {
        path = append(path, root.Val)
        tmp := make([]int, len(path))
        copy(tmp, path)
        *paths = append(*paths, tmp)
        return
    }
    
    // process current logic
    path = append(path, root.Val)
    dfs(root,paths,path,targetSum - root.Val)
    dfs(root,paths,path,targetSum - root.Val)
}

执行 curreng logic 的时候错了，应该是 root.Left 和 root.Right 的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    paths := [][]int{}
    path := []int{}
    
    dfs(root,&paths,path,targetSum)
    
    return paths
}

func dfs(root *TreeNode,paths *[][]int,path []int,targetSum int) {
    // terminator
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && root.Val == targetSum {
        path = append(path, root.Val)
        tmp := make([]int, len(path))
        copy(tmp, path)
        *paths = append(*paths, tmp)
        return
    }
    
    // process current logic
    path = append(path, root.Val)
    dfs(root.Left,paths,path,targetSum - root.Val)
    dfs(root.Right,paths,path,targetSum - root.Val)
}

3. 看题解

看BFS的题解，看怎么存储 queue，以及怎么找到值以后找到路径信息：
定义 pair，然后对队列进行操作
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Pair struct {
    node *TreeNode
    left int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }
    parent := map[*TreeNode]*TreeNode{}
    paths := [][]int{}
    
    queue := []Pair{{root, targetSum}}
    
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        node := p.node
        left := p.left - node.Val
        
        if node.Left == nil && node.Right == nil && left == 0 {
            paths = append(paths, getPath(node, parent))
        }else {
            if node.Left != nil {
                parent[node.Left] = node
                queue = append(queue, Pair{node.Left, left})
            }
            
            if node.Right != nil {
                parent[node.Right] = node
                queue = append(queue, Pair{node.Right, left})
            }
        }
    }
    
    return paths
}

func getPath(node *TreeNode,parent map[*TreeNode]*TreeNode) []int {
    path := []int{}
    
    for ; node != nil; node = parent[node] {
        path = append(path, node.Val)
    }
    
    for i,j := 0,len(path) - 1;i < j;i++ {
        path[i],path[j] = path[j],path[i]
        j--
    }
    
    return path
}

4. 复杂度分析
时间复杂度：O(N), N是节点个数
空间复杂度：O(N), 递归调用栈空间和开辟队列大小

5. 总结
  5.1: 这道题应该是第3边做了，还是磕磕绊绊，做过不等于会，要经常的反复去练习
  5.2: BFS的时候，定义 Pair 和 找到 parent 路径这块代码不会写，要多写，多练，多思考
