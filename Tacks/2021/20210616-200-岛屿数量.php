<?php
/*
 * @Descripttion: 周三搞
 * @Author: tacks321@qq.com
 * @Date: 2021-06-16 17:11:07
 * @LastEditTime: 2021-06-16 19:35:31
 */

/*
 * @lc app=leetcode.cn id=200 lang=php
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (54.06%)
 * Likes:    1194
 * Dislikes: 0
 * Total Accepted:    264.4K
 * Total Submissions: 488.6K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 
 * 此外，你可以假设该网格的四条边均被水包围。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 
 * grid[i][j] 的值为 '0' 或 '1'
 * 
 * 
 */

// @lc code=start
class Solution {


    // 四个方向
    private static $DIRECTIONS = [[-1,0], [0,1], [1,0], [0,-1]]; // 上右下左
    
    // 网格大小
    private $rows;
    private $cols;

    private $grid;    // 网格
    private $visited; // 访问过的点

    /**
     * @param String[][] $grid
     * @return Integer
     */
    // DFS 深搜
    // 执行用时：68 ms, 在所有 PHP 提交中击败了24.24%的用户
    // 内存消耗：22.9 MB, 在所有 PHP 提交中击败了36.36%的用户
    function numIslands($grid) {
        // 赋值边界
        $this->rows = count($grid);
        if(!$this->rows) {
            return 0;
        }
        $this->cols = count($grid[0]);
        $this->grid = $grid;
        $this->visited = [];

        // 初始化 visited数组
        $this->visited = array_fill(0, $this->rows,[]);
        for($i=0; $i < $this->rows; $i++){
            $this->visited[$i] = array_fill(0,$this->cols,0);
        }

         
        // 岛屿数量
        $count = 0;

        for ($i=0; $i < $this->rows; $i++) { 
            for ($j=0; $j < $this->cols; $j++) { 
                // 如果岛屿中的陆地，没有被访问过，就进行深搜
                if($this->grid[$i][$j] == 1 && !$this->visited[$i][$j]) {
                    $this->_dfs($i, $j);
                    // 每次深搜过一片区域进行 记录
                    $count++;
                }
            }
        }

        return $count;

    }


    /**
     * DFS 深搜遍历
     *
     * @param [int] $i 坐标j
     * @param [int] $j 坐标i
     */
    private function _dfs($i, $j) {
        // 访问过进行标记
        $this->visited[$i][$j] = true;
        // 四个节点进行遍历
        for ($k=0; $k<4; $k++) { 
            $newX = $i + self::$DIRECTIONS[$k][0];
            $newY = $j + self::$DIRECTIONS[$k][1];

            // 判断当前节点陆地是否没有越界，并且没有被访问过
            if($this->_isInArea($newX, $newY) 
                && $this->grid[$newX][$newY] == 1
                && !$this->visited[$newX][$newY]) {
                    // 递归处理
                    $this->_dfs($newX, $newY);
            }
        }
    }


    /**
     * 判断是否在网格中
     *
     * @date  2021-06-16 18:13:02
     * @param [int] $x
     * @param [int] $y
     */
    private function _isInArea($x, $y) {
        return $x >= 0 && $x < $this->rows && $y >= 0 && $y < $this->cols;
    }
}



class Solution2 {


    // 四个方向
    private static $DIRECTIONS = [[-1,0], [0,1], [1,0], [0,-1]]; // 上右下左

    // 网格大小
    private $rows;
    private $cols;

    private $grid;    // 网格
    private $visited; // 访问过的点

    /**
     * @param String[][] $grid
     * @return Integer
     */
    // BFS 广搜
    // 执行用时：72 ms, 在所有 PHP 提交中击败了19.70%的用户
    // 内存消耗：22.7 MB, 在所有 PHP 提交中击败了71.21%的用户
    function numIslands($grid) {
        // 赋值边界
        $this->rows = count($grid);
        if(!$this->rows) {
            return 0;
        }
        $this->cols = count($grid[0]);
        $this->grid = $grid;
        $this->visited = [];

        // 初始化 visited数组
        $this->visited = array_fill(0, $this->rows,[]);
        for($i=0; $i < $this->rows; $i++){
            $this->visited[$i] = array_fill(0,$this->cols,0);
        }

 
        // 岛屿数量
        $count = 0;

        for ($i=0; $i < $this->rows; $i++) { 
            for ($j=0; $j < $this->cols; $j++) { 
                // 如果岛屿中的陆地，没有被访问过，就进行深搜
                if($this->grid[$i][$j] == 1 && !$this->visited[$i][$j]) {
                    // 广搜
                    $this->_bfs($i, $j);
                    // 每次广搜过一片区域进行 记录
                    $count++;
                }
            }
        }

        return $count;

    }


    /**
     * BFS 广搜遍历
     *
     * @param [int] $i 坐标j
     * @param [int] $j 坐标i
     */
    private function _bfs($i, $j) {
        $queue = []; // 队列

        array_push($queue, $i * $this->cols + $j); // 节点入队
        
        // 【注意】访问过进行标记
        $this->visited[$i][$j] = true;

        // 队列遍历
        while(count($queue)) {
            $item = array_shift($queue);  // 队列 先进先出 移除头部节点

            // 获取当前坐标 【注意PHP的除法运算】
            $itemX = intval($item / $this->cols);
            $itemY = intval($item % $this->cols);

            // 四个节点进行遍历
            for ($k=0; $k<4; $k++) { 
                $newX = $itemX + self::$DIRECTIONS[$k][0];
                $newY = $itemY + self::$DIRECTIONS[$k][1];

                // 判断当前节点陆地是否没有越界，并且没有被访问过
                if($this->_isInArea($newX, $newY) 
                    && $this->grid[$newX][$newY] == 1
                    && !$this->visited[$newX][$newY]) {

                        array_push($queue, $newX * $this->cols + $newY); // 压入队列

                        $this->visited[$newX][$newY] = true; // !!! 【注意】  只要入队就进行标记
                }
            }
        }
     
    }


    /**
     * 判断是否在网格中
     *
     * @date  2021-06-16 18:13:02
     * @param [int] $x
     * @param [int] $y
     */
    private function _isInArea($x, $y) {
        return $x >= 0 && $x < $this->rows && $y >= 0 && $y < $this->cols;
    }
}
// @lc code=end


$Obj = new Solution2();
var_dump($Obj->numIslands([["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]));

/*
【问题】
    岛屿问题
        1为陆地
        0为水域
    
    网格问题 => 利用二维数组来构建岛屿 => 定义一个m*n的二维数组
    二叉树的DFS/BFS => 网格上的DFS/BFS

   
【参考】
https://leetcode-cn.com/problems/number-of-islands/solution/dfs-bfs-bing-cha-ji-python-dai-ma-java-dai-ma-by-l/ 


【解析】

    Flood Fill 算法
        从一个区域中提取若干个连通的点与其他相邻区域区分开（或分别染成不同颜色）的经典算法。
        因为其思路类似洪水从一个区域扩散到所有能到达的区域而得名。
    具体实现
        DFS、BFS 通过一个点，不断找到其他的联通点，并将相连的格子进行标记，形成一个岛屿


1、DFS深搜
    DFS基本结构
        function dfs($root) {
            // base case 
            if($root == null) {
                return 0;
            }
            // 访问左右节点
            $this->dfs($root->left);
            $this->dfs($root->right);
        }
    
    网格的相邻节点
        当前节点(r, c)
        四个相邻节点 (r-1,c) (r,c+1) (r+1,c) (r,c-1)
    
    网格的DFS
        1、base case
            判断是否需要继续遍历，也就是是否会越界超出二维数组
        2、访问节点
            访问 上、右、下、左 四个节点
        3、特别处理
            由于网格结构本质为图，对每个节点的上右下左遍历一次，则会遍历重复，导致兜圈子
            所以需要进行标记
                visited进行标记

    复杂度
        时间复杂度：O(M*N)，其中 M 和 N 分别为行数和列数。
        空间复杂度：O(M*N)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 M*N。


2、BFS广搜
    BFS基本结构
        while(count($queue)) {
            $node = array_shift($queue);  // 队列 先进先出 移除头部节点
            
            if($node->left != null ) {
                $queue[] = $node->left;   // 入队
            }
            if($node->right != null) {
                $queue[] = $node->right;  // 入队
            }
        }
    借助队列
        注意标记的时机
            特别注意：在放入队列以后，要马上标记成已经访问过。
            只要进入了队列，迟早都会遍历到它
        为什么不是出队列的时候标记？
            如果是出队列的时候再标记，会造成很多重复的结点进入队列，可能会造成重复的操作

    复杂度
        时间复杂度：O(M*N)，其中 M 和 N 分别为行数和列数。
        空间复杂度：O(min(M,N))，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 min(M,N)。

【总结】
    1、掌握DFS、BFS的基础思想，写出自己的模板。
    2、根据题意，在原有的DFS\BFS上进行添加代码，核心仍然不变。
    3、细节仍需注意，base case 什么时候可以访问下面的节点、 visited如何防止遍历重复 、PHP除法运算注意int处理
 
            
*/