<?php
/*
 * @lc app=leetcode.cn id=1267 lang=php
 *
 * [1267] 统计参与通信的服务器
 *
 * https://leetcode-cn.com/problems/count-servers-that-communicate/description/
 *
 * algorithms
 * Medium (60.20%)
 * Likes:    22
 * Dislikes: 0
 * Total Accepted:    5.4K
 * Total Submissions: 8.9K
 * Testcase Example:  '[[1,0],[0,1]]'
 *
 * 这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。
 * 
 * 如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。
 * 
 * 请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：grid = [[1,0],[0,1]]
 * 输出：0
 * 解释：没有一台服务器能与其他服务器进行通信。
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：grid = [[1,0],[1,1]]
 * 输出：3
 * 解释：所有这些服务器都至少可以与一台别的服务器进行通信。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
 * 输出：4
 * 解释：第一行的两台服务器互相通信，第三列的两台服务器互相通信，但右下角的服务器无法与其他服务器通信。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m <= 250
 * 1 <= n <= 250
 * grid[i][j] == 0 or 1
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
     // 计数法
     // 执行用时：204 ms, 在所有 PHP 提交中击败了100.00%的用户
     // 内存消耗：21.8 MB, 在所有 PHP 提交中击败了100.00%的用户
    function countServers1($grid) {
        $count_m = []; // 统计行位置上 值=1的 可以通信的服务器
        $count_n = []; // 统计列位置上 值=1的 可以通信的服务器
        foreach($grid as $i=>$item) {
            foreach($item as $j=>$val) {
                if($val == 1) { // 计数
                    $count_m[$i]++;
                    $count_n[$j]++;
                } 
            }
        }
        $res = 0;
        foreach($grid as $i=>$item) {
            foreach($item as $j=>$val) {
                // 判断 只要在i行或者j列，满足超过一台服务器即可进行通信。
                if($val == 1 && ($count_m[$i] > 1 || $count_n[$j] >1) ) {
                   $res++;
                } 
            }
        }
        return $res;
    }

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function countServers($grid) {
        $sum = 0;
        foreach($grid as $i=>$item) {
            foreach($item as $j=>$val) {
                if($val == 1) {
                   $res = $this->dfs($grid, $i, $j);
                   var_dump($res);
                   if($res > 1) {
                       $sum += $res;
                   }
                } 
            }
        }
        return $sum;
    }

    // 深搜
    function dfs($grid, $x, $y) {
        $row = count($grid);   // 行
        $col = count($grid[0]);// 列
        if($x < 0 || $x > $row - 1 || $y < 0 || $y > $col - 1) {
            return 0;
        }
        if($grid[$x][$y] != 1) {
            return 0;
        }
        $grid[$x][$y] = -1; //标记扫描过的
        $sum = 1;

        //搜索整列
        for($i=0; $i<$row; $i++) {
            $sum += $this->dfs($grid, $i, $y);
        }
        // 搜索整行
        for($i=0; $i<$col; $i++) {
            $sum += $this->dfs($grid, $x, $i);
        }
        return $sum;
    }
   
}
// @lc code=end

$obj = new Solution();
$obj->countServers([[1,0],[1,1]]);
// @tacks think=start
/*
题意 meaning 
  
    对于某个服务器上，坐标为(x,y) ，也就是说在x上至少有两台或者y上有两台就可以满足通信，否则不行。

关键 key 
     
    “通信条件”

想法 idea

【1】计数法
   
    遍历整个数组统计行和列的服务器个数，
    然后再遍历每个服务器，判断它的行和列的服务器个数是否大于1，进行计数。

    时间复杂度，双层遍历O(N*M)
    需要维护两个数组记录服务器的个数，空间复杂度O(N+M)

【2】深搜

    看了小马姐的题解，整行整列进行搜索。待实现
  
    

*/
// @tacks think=end
