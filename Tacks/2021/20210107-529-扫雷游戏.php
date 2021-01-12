<?php
/*
 * @Descripttion: 今天差点都不想做题了，这道看起来好复杂，情绪也不好，哎
 * @Author: tacks321@qq.com
 * @Date: 2021-01-07 19:54:12
 * @LastEditTime: 2021-01-07 20:46:22
 */

/*
 * @lc app=leetcode.cn id=529 lang=php
 *
 * [529] 扫雷游戏
 *
 * https://leetcode-cn.com/problems/minesweeper/description/
 *
 * algorithms
 * Medium (65.40%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    31.5K
 * Total Submissions: 48.2K
 * Testcase Example:  '[["E","E","E","E","E"],["E","E","M","E","E"],["E","E","E","E","E"],["E","E","E","E","E"]]\n' +
  '[3,0]'
 *
 * 让我们一起来玩扫雷游戏！
 * 
 * 给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B'
 * 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X'
 * 则表示一个已挖出的地雷。
 * 
 * 现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
 * 
 * 
 * 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
 * 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
 * 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
 * 如果在此次点击中，若无更多方块可被揭露，则返回面板。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入: 
 * 
 * [['E', 'E', 'E', 'E', 'E'],
 * ⁠['E', 'E', 'M', 'E', 'E'],
 * ⁠['E', 'E', 'E', 'E', 'E'],
 * ⁠['E', 'E', 'E', 'E', 'E']]
 * 
 * Click : [3,0]
 * 
 * 输出: 
 * 
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'M', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 * 
 * 解释:
 * 
 * 
 * 
 * 示例 2：
 * 
 * 输入: 
 * 
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'M', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 * 
 * Click : [1,2]
 * 
 * 输出: 
 * 
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'X', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 * 
 * 解释:
 * 
 * 
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 输入矩阵的宽和高的范围为 [1,50]。
 * 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
 * 输入面板不会是游戏结束的状态（即有地雷已被挖出）。
 * 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
 * 
 * 
 */

// @lc code=start
class Solution {


    protected $direction; // 定义八个方向
    /**
     * @param String[][] $board
     * @param Integer[] $click
     * @return String[][]
     */
    // 执行用时：56 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：17.2 MB, 在所有 PHP 提交中击败了81.25%的用户
    function updateBoard($board, $click){
        // 边界判断
        if ($board == null || count($board) < 1 || $board[0] == null || count($board[0]) < 1 || $click == null || count($click) != 2) {
            return $board;
        }
        $m = count($board);
        $n = count($board[0]);
        $x = $click[1];
        $y = $click[0];
        // 边界判断
        if ($x >= $n || $x < 0 || $y >= $m || $y < 0) {
            return $board;
        }
        // 如果直接点到炸弹爆炸 Over
        if ($board[$y][$x] == 'M') {
            $board[$y][$x] = 'X';
            return $board;
        }
        // 初始化 在原坐标的基础上依次和如下的值进行8次相加
        $this->direction = [
            [-1, 0],
            [-1, 1],
            [0, 1],
            [1, 1],
            [1, 0],
            [1, -1],
            [0, -1],
            [-1, -1],
        ];
        $this->process($board, $x, $y);
        return $board;
    }
 
    function process(&$board, $x, $y)
    {
        $m = count($board);
        $n = count($board[0]);
        if($x<0 || $x>=$n || $y<0 || $y>=$m)
        {
            return;
        }
        if($board[$y][$x] == 'E') {
            $board[$y][$x] = 'B';
            $count = $this->boom($board, $x, $y);
            
            if($count != 0) {
                $board[$y][$x] = chr($count + ord('0'));
            } else {
                foreach ($this->direction as $item) {
                    $spaceX = $x + $item[1];
                    $spaceY = $y + $item[0];
                    $this->process($board, $spaceX, $spaceY);
                }
            }
        }

    }
    
    // 查找一个格子周围的雷的数目
    function boom($board, $x, $y){
        $m = count($board);
        $n = count($board[0]);
        $count = 0;
        foreach ($this->direction as $item) {
            $boomX = $x + $item[1];
            $boomY = $y + $item[0];
            // 越界了，继续找周围其他格子
            if ($boomX < 0 || $boomX >= $n || $boomY < 0 || $boomY >= $m) {
                continue;
            }
            // 找到一颗雷
            if ($board[$boomY][$boomX] == 'M') {
                $count++;
            }
        }
        return $count;
    }

}
// @lc code=end

/*
采用模拟办法
    当前点击的如果是地雷 M => 直接将对应的值改为 X
    当前点击的如果是空方块 => 统计周围相邻的方块地雷的数量cnt 也就是M的数量
        如果cnt = 0 ,我们就将其改为B，然后递归的处理周围八个未挖的方块，递归终止条件是：没有更多方块被揭露
        否则执行将其修改为数字
    
通过点击一个点的位置，逐渐向四面八方扩展，采用搜索方式。

【深度搜索】DFS

最后还是一步一步看看题解，理解整体思路，模拟一下游戏过程


遇到这种应用题，总是内心怕怕的，但其实读懂了，慢慢也是能抠出来的


最近状态不太好，刷题还是要坚持
*/