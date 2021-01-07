// 529. 扫雷游戏
// 让我们一起来玩扫雷游戏！

// 给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。

// 现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

// 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
// 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
// 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
// 如果在此次点击中，若无更多方块可被揭露，则返回面板。
 

// 示例 1：

// 输入: 

// [['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'M', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E']]

// Click : [3,0]

// 输出: 

// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]

// 解释:

// 示例 2：

// 输入: 

// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]

// Click : [1,2]

// 输出: 

// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'X', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]

// 解释:

 

// 注意：

// 输入矩阵的宽和高的范围为 [1,50]。
// 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
// 输入面板不会是游戏结束的状态（即有地雷已被挖出）。
// 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。


/*
 * @lc app=leetcode.cn id=529 lang=javascript
 *
 * [529] 扫雷游戏
 */

// @lc code=start
/**
 * @param {character[][]} board
 * @param {number[]} click
 * @return {character[][]}
 */
var updateBoard = function (board, click) {
    let direction = [[1, 1], [1, -1], [1, 0], [-1, 1], [-1, 0], [-1, -1], [0, 1], [0, -1]]
    let overBorder = (x, y) => {
        return (x < 0 || x >= board.length || y < 0 || y >= board[0].length)
    }
    let dfs = (x, y) => {
        if (overBorder(x, y) || board[x][y] != 'E') {
            return;
        }
        let count = 0
        //记录周围雷的个数
        for (let i = 0; i < direction.length; i++) {
            let goX = x + direction[i][0], goY = y + direction[i][1]
            if (!overBorder(goX, goY) && board[goX][goY] == 'M') {
                count++
            }
        }
        if (!count) {//周围没雷则改为'B',且递归八个方向
            board[x][y] = 'B'
            for (let i = 0; i < direction.length; i++) {
                let goX = x + direction[i][0], goY = y + direction[i][1]
                dfs(goX, goY)
            }
        }
        else {
            //有雷改为值
            board[x][y] = count + ''
        }
    }
    let x = click[0], y = click[1]
    if (board[x][y] == 'M') {
        board[x][y] = 'X'
    }
    else {
        dfs(x, y)
    }
    return board
};

/* 
    递归。
    主要逻辑:1.点击的是否是地雷，是，直接改为'X'且返回，不是则开始递归之路
            2.递归过程：1）首先出边界了或者当前位置不是'E'，则递归结束。
                      2）剩余都是'E'的处理了：
                         然后看周围有没有相邻的地雷。则需要八个方向确认一下，有则记录相邻个数，并修改值为个数
                         无则值修改为'B',同时相邻的方块也应该递归，再来一次八个方向的递归
    时间复杂度：O(N*M)
    每个位置都会被扫描一次
    空间复杂度：O(N*M)
    递归节点所占用空间


*/


//看了题解，发现模拟思路基本是一致的，而队列迭代的处理就是在上边递归的位置，改为push入队
//多多练习吧，但感觉类似这种题肯定是递归写出来才去模拟队列，因为在这种结构队列的逻辑性太弱了。


var updateBoard = function (board, click) {
    let direction = [[1, 1], [1, -1], [1, 0], [-1, 1], [-1, 0], [-1, -1], [0, 1], [0, -1]]
    let overBorder = (x, y) => {
        return (x < 0 || x >= board.length || y < 0 || y >= board[0].length)
    }
    //相邻adjacent
    let bfs = (x, y) => {
        if (overBorder(x, y) || board[x][y] != 'E') {
            return;
        }
        let queue = [[x, y]]
        while (queue.length) {
            let [curX,curY]=queue.shift()
            let count = 0
            for (let i = 0; i < direction.length; i++) {
                let goX = curX + direction[i][0], goY = curY + direction[i][1]
                if (!overBorder(goX, goY) && board[goX][goY] == 'M') {
                    count++
                }
            }
            if (!count) {
                board[curX][curY] = 'B'
                for (let i = 0; i < direction.length; i++) {
                    let goX = curX+ direction[i][0], goY = curY + direction[i][1]
                    //[goX,goY]入队：首先不出界，且方块位置为'E'
                    if(!overBorder(goX,goY) && board[goX][goY] == 'E')
                    {   
                        board[goX][goY]='B'
                        queue.push([goX,goY])
                    }
                }
            }
            else {
                board[curX][curY] = count + ''
            }
        }
    }
    let x = click[0], y = click[1]
    if (board[x][y] == 'M') {
        board[x][y] = 'X'
    }
    else {
        bfs(x, y)
    }
    return board
};
