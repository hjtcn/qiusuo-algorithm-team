你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
 
提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

1. Clearfication:
搜索问题：
        BFS:
            只知道是用BFS去解决，怎么写呢？
            印象中是8联通
            helper(grid,)

2. Coding:
func numIslands(grid [][]byte) int {
}
func helper(grid [][]byte,i,j int) {
    helper(grids,i - 1,j-1)
    helper(grid,i, j - 1)
    helper(grid,i-1,j+1)
    helepr(grid,i-1,j)
    helper(grid,i+1,j)
    helper(grid,i+1,j-1)
    helper(gird,i,j+1)
    helepr(grid,i+1,j+1)
}

3. 看题解：
func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                count++
            }
        }
    }
    return count
}

func dfs(grid [][]byte, i, j int) {
    if (i < 0 || j < 0) || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j-1)
    dfs(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
    q := make([][2]int, 0, len(grid[0])) // queue
    directions := [][]int{{1, 0},{0, 1}, {-1,0}, {0,-1}} // directions to move from a '1' block
    islands := 0    
    
    // loop over all coordinates
    // when we hit a 1, we BFS and mark visited nodes as 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                q = append(q, [2]int{i, j})
                islands++
                grid[i][j] = '0' // visited
                for len(q) > 0 {
                    // grab from queue
                    row, col := q[0][0], q[0][1]
                    // pop from queue
                    q = q[1:]
                    // search nearby coords
                    for _, d := range directions {
                        nextRow, nextCol := row + d[0], col + d[1]
                        // check boundaries
                        if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
                            continue
                        }
                        // if we find another '1', add coords to queue and set to visited '0'
                        if grid[nextRow][nextCol] == '1' {
                            q = append(q, [2]int{nextRow, nextCol})
                            grid[nextRow][nextCol] = '0'
                        }
                    }
                }
            }
        }
    }
    return islands
}

4. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n)

5. 总结：
5.1: 感觉脑子转不动了，不知道是加班太多，还是太忙了
