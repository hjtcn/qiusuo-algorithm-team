/*
给定一个包含了一些 0 和 1 的非空二维数组 grid 。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
示例 2:
[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。
注意: 给定的矩阵grid 的长度和宽度都不超过 50。
*/

1. Clearfiction:
1.1 visited 数组记录是否访问过该节点
1.2 记录结果，判断当前遍历过的结果是否大于历史结果
细节点：
如何组织代码更容易阅读
如何判断是否访问过
如何更新遍历的结果

func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    visited := make([][]int,m)
    
    for i := 0;i < m;i++ {
        visited[i] = make([]int, n)
    }
    ret := 0
    
    for i := 0;i < m;i++ {
        for j := 0;j < n;j++ {
            if grid[i][j] == 1 && visited[i][j] == 0 {
                current := 0
                dfs(grid,visited,i,j,m,n,&current,&ret)
                
                if current > ret {
                    ret = current
                }
            }
        }   
    }
   
    return ret
}
func dfs(grid [][]int,visited [][]int, i,j,m,n int,current *int,ret *int) {
    if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 && visited[i][j] == 0 {
         visited[i][j] = 1
         *current = *current + 1
        
        dfs(grid,visited,i - 1,j,m,n,current,ret)
        dfs(grid,visited,i + 1,j,m,n,current,ret)
        dfs(grid,visited,i,j - 1,m,n,current,ret)
        dfs(grid,visited,i,j + 1,m,n,current,ret)
    }else {
        return
    }
}

DFS自己可以写出来了，BFS的话节点怎么变化，怎么放入到队列中，然后怎么取出来，自己并没有想清楚了

func maxAreaOfIsland(grid [][]int) int {
    rowNum := len(grid)
    if rowNum == 0 {
        return 0
    }
    colNum := len(grid[0])
    max := 0 
    for r, v := range grid {
        for c, v1 := range v {
            if v1 == 0 {
                continue
            }
            connected, curArea := [][]int{{r, c}}, 0
            grid[r][c] = 0 
            for len(connected) != 0 {
                top := connected[0] 
                connected = connected[1:]
                curArea++
                row, col := top[0], top[1] 
                if row-1 >= 0 && grid[row-1][col] == 1 {
                    grid[row-1][col] = 0
                    connected = append(connected, []int{row - 1, col})
                }
                if row+1 < rowNum && grid[row+1][col] == 1 { 
                    grid[row+1][col] = 0
                    connected = append(connected, []int{row + 1, col})
                }
                if col-1 >= 0 && grid[row][col-1] == 1 { 
                    grid[row][col-1] = 0
                    connected = append(connected, []int{row, col - 1})
                }
                if col+1 < colNum && grid[row][col+1] == 1 { 
                    grid[row][col+1] = 0
                    connected = append(connected, []int{row, col + 1})
                }
            }
            if curArea > max {
                max = curArea
            }
        }
    }
    return max
}

复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n) 递归栈空间和存储队列长度

总结：
1. 和自己斗争是一个长期和持久的过程

2. 不要害怕困难，一点点分析
