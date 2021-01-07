/*
让我们一起来玩扫雷游戏！
给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。
现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
1. 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
2. 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
3. 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
4. 如果在此次点击中，若无更多方块可被揭露，则返回面板。
 
示例 1：
输入: 
[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]
Click : [3,0]
输出: 
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]
解释:
示例 2：
输入: 
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]
Click : [1,2]
输出: 
[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'X', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]
解释:
注意：
输入矩阵的宽和高的范围为 [1,50]。
点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
输入面板不会是游戏结束的状态（即有地雷已被挖出）。
简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
*/

1. Clearfication:
看着题目很吓人，我们一点点的来分析：
主要是下面这4点：
1. 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
2. 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
3. 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
4. 如果在此次点击中，若无更多方块可被揭露，则返回面板。

1. 如果board[i][j] == 'X' 将它周围的都改为 '1' return false
2. 如果没有相邻地雷的空方块（'E') 被挖出，修改它为('B'),并且所有和其相邻的未挖出方块都应该被递归
看着像是if else,然后递归，但是我脑子有点疼，哈哈哈哈哈哈
3. 如果一个至少与一个地雷相邻的空方块E被挖出，修改它为数字1到8，表示相邻地雷的数量
if E 与地雷相邻，修改它为数字 1 到 8
4. 如果在此次点击中，没有更多方块可被揭露，返回面板

var dirX = []int{0,1,0,-1,1,1,-1,-1}
var dirY = []int{1,0,-1,0,1,-1,1,-1}
func updateBoard(board [][]byte,click []int) [][]byte {
    x,y := click[0],click[1]
    
    if board[x][y] == 'M' {
        board[x][y] = 'X'
    }else {
        dfs(board, x,y)
    }
    return board
}
func dfs(board [][]byte,x,y int) {
    cnt := 0
    for i := 0;i < 8;i++ {
        tx,ty := x + dirX[i],y + dirY[i]
        
        if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
            continue
        }
        
        if board[tx][ty] == 'M' {
            cnt++
        }
    }
    
    if cnt > 0 {
        board[x][y] = byte(cnt + '0')
    }else {
        board[x][y] = 'B'
        for i := 0;i < 8;i++ {
            tx,ty := x + dirX[i],y + dirY[i]
            
            if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
                continue
            }
            
            dfs(board, tx,ty)
        }
    }
}

BFS:
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}
func updateBoard(board [][]byte, click []int) [][]byte {
    x, y := click[0], click[1]
    if board[x][y] == 'M' {
        board[x][y] = 'X'
    } else {
        bfs(board, x, y)
    }
    return board
}
func bfs(board [][]byte, sx, sy int) {
    queue := [][]int{}
    vis := make([][]bool, len(board))
    for i := 0; i < len(vis); i++ {
        vis[i] = make([]bool, len(board[0]))
    }
    queue = append(queue, []int{sx, sy})
    vis[sx][sy] = true
    for i := 0; i < len(queue); i++ {
        cnt, x, y := 0, queue[i][0], queue[i][1]
        for i := 0; i < 8; i++ {
        tx, ty := x + dirX[i], y + dirY[i]
            if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
                continue
            }
            if board[tx][ty] == 'M' {
                cnt++
            }
        }
        if cnt > 0 {
            board[x][y] = byte(cnt + '0')
        } else {
            board[x][y] = 'B'
            for i := 0; i < 8; i++ {
                tx, ty := x + dirX[i], y + dirY[i]
                if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' || vis[tx][ty] {
                    continue
                }
                queue = append(queue, []int{tx, ty})
                vis[tx][ty] = true
            }
        }
    }
}

复杂度分析：
时间复杂度：O（nm)
空间复杂度：O（nm)

总结：
1. 题目没理解清楚

2. 同样是if else for loop,但是写出来的代码并不是if else for loop 这么简单。。。
