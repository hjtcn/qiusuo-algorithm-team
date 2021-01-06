/*
班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
示例 1：
输入：
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出：2 
解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回 2 。
示例 2：
输入：
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出：1
解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1 。
提示：
1 <= N <= 200
M[i][i] == 1
M[i][j] == M[j][i]
*/

1. Clearfication:

遇到矩阵，我在想可以用昨天的方式进行做么，直接遍历，然后DFS或者BFS将周围节点状态改变，记录状态这样可以么
目前我想的是不可以的，举例说明，如果1和4是朋友，那么[0,3]和[3,0]都为1,他们是1个朋友圈，如果像水域大小那道题一样的话，它会计算2个，所以感觉还是怪怪的
想了一会没想出来看了题解：
看了题解以后看出来，他还是将问题抽象了出来，抽象成了求图中的连通块的数量

func dfs(M[][]int, visited []int,i int) {
    length := len(M)
    for j := 0;j < length;j++ {
        if M[i][j] == 1 && visited[j] == 0 {
            visited[j] = 1
            dfs(M,visited,j)
        }
    }
}

func findCircleNum(M [][]int) int {
     length := len(M)
     visited := make([]int, length)
     count := 0
     for i := 0;i < length;i++ {
         if visited[i] == 0 {
             dfs(M,visited,i)
             count++
         }
     }
     return count
}

还引起的思考是初始化切片的时候，visitedd 切片变量传递的时候，它的值是如何变化和进行的

func findCircleNum(M [][]int) int {
     length := len(M)
     visited := make([]int, length)
     count := 0
    
     queue := []int{}
   
     for i := 0;i < length;i++ {
         if visited[i] == 0 {
             queue = append(queue,i)
             for len(queue) > 0 {
                 s := queue[0]
                 queue = queue[1:]
                 visited[s] = 1
                 for j := 0;j < length;j++ {
                     if M[s][j] == 1 && visited[j] == 0 {
                         queue = append(queue,j)
                     }
                 }
             }
             count++
         }
     }
     return count
}

union-find 并查集

find(parent []int,i int) {
    if parent[i] == -1 {
        return i
    }
    return find(parent, parent[i])
}

func union(parent []int,x,y int) {
    xset := find(parent,x)
    yset := find(parent,y)
    if xset != yset {
        parent[xset] = yset
    }
}

func findCircleNum(M [][]int) int {
    length := len(M)
    parent := make([]int, length)
    
    for i := 0;i < length;i++ {
        parent[i] = -1
    }
    
    for i := 0;i < length;i++ {
        for j := 0;j < length;j++ {
            if M[i][j] == 1 && i != j {
                union(parent,i,j)
            }
        }
    }
    
    count := 0
    for i := 0;i < length;i++ {
        if parent[i] == -1 {
            count++
        }
    }
    
    return count
}  

复杂度分析：
时间复杂度：DFS/BFS: O(n*n) 要遍历每个节点，union-find (并查集) O(n*n*n) 遍历矩阵O(n*n) 构建并查集最坏时间复杂度为O(n),所以是O(n*n*n)
空间复杂度：O(n) 递归栈调用空间，存储队列节点，并查集parent 数组

总结：
1. 搞清楚思路，然后一点点的看代码

2. 遇到不会的不要害怕，一点点看就好了
