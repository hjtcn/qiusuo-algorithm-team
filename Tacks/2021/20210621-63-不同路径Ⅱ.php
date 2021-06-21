<?php
/*
 * @Descripttion: 
 * @Author: tacks321@qq.com
 * @Date: 2021-06-21 15:24:16
 * @LastEditTime: 2021-06-21 18:17:39
 */


/*
 * @lc app=leetcode.cn id=63 lang=php
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (38.37%)
 * Likes:    572
 * Dislikes: 0
 * Total Accepted:    152.8K
 * Total Submissions: 397.3K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 * 
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 * 
 * 
 * 
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
 * 输出：2
 * 解释：
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：obstacleGrid = [[0,1],[0,0]]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == obstacleGrid.length
 * n == obstacleGrid[i].length
 * 1 
 * obstacleGrid[i][j] 为 0 或 1
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    // 动态规划
    // 执行用时：8 ms, 在所有 PHP 提交中击败了97.56%的用户
    // 内存消耗：15.8 MB, 在所有 PHP 提交中击败了14.64%的用户
    function uniquePathsWithObstacles($obstacleGrid) {
        $m = count($obstacleGrid);
        $n = count($obstacleGrid[0]);
        // base case
        if(!$m || !$n || $obstacleGrid[0][0] == 1 || $obstacleGrid[$m-1][$n-1] == 1) {
            return 0;
        }

        $dp = [];
        // 初始化都为0
        $dp = array_fill(0, $m, array_fill(0, $n, 0));
        
        $dp[0][0] = 1; // 终点就是出发点

        // 第一列的case
        for ($i=0; $i < $m && $obstacleGrid[$i][0] == 0; $i++) { 
            // 初始化左边界，遇到障碍就退出，比如 $obstacleGrid[0][1] == 1 退出了，$dp[0][$i]后续仍保持 0 
            $dp[$i][0] = 1;
        }
        // 第一行的case
        for ($j=0; $j < $n && $obstacleGrid[0][$j] == 0; $j++) { 
            // 初始化上边界
            $dp[0][$j] = 1;
        }

        // 动态转移 dp[i][j] = dp[i-1][j] + dp[i][j-1] 
        for ($i=1; $i < $m; $i++) { 
            for ($j=1; $j < $n; $j++) { 
                if($obstacleGrid[$i][$j] == 0) {
                    $dp[$i][$j] = $dp[$i-1][$j] + $dp[$i][$j-1];
                }
            }
        }

        return $dp[$m-1][$n-1];  // 到达(m-1,n-1) 终点的路径数
    }
}

class Solution2 {

    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    // 动态规划 一维数组
    // 执行用时：12 ms, 在所有 PHP 提交中击败了75.61%的用户
    // 内存消耗：15.9 MB, 在所有 PHP 提交中击败了12.20%的用户
    function uniquePathsWithObstacles($obstacleGrid) {
        $m = count($obstacleGrid);
        $n = count($obstacleGrid[0]);
        // base case
        if(!$m || !$n || $obstacleGrid[0][0] == 1 || $obstacleGrid[$m-1][$n-1] == 1) {
            return 0;
        }

        $dp = [0];
        $dp = array_pad($dp, $m, 0); // 填充一下
        $dp[0] = 1;
        
        // 动态转移 dp[i][j] = dp[i-1][j] + dp[i][j-1] 
        for ($i=0; $i < $m; $i++) { 
            for ($j=0; $j < $n; $j++) { 
                if($obstacleGrid[$i][$j] == 1) {
                    $dp[$j] = 0;
                }else{
                    $dp[$j] += $dp[$j-1];
                }
            }
        }

        return $dp[$n-1];  
    }
}

// @lc code=end

/*
【题目】

动态规划的题目分为两大类（存在一定的递推性质）
    一种是求最优解类，典型问题是背包问题 （最优子结构/即当前问题的最优解取决于子问题的最优解）
    另一种就是计数类，如统计方案数的问题 （当前问题的方案数取决于子问题的方案数）

    
【解析】

动态规划

①状态定义
    我们用 f(i, j) 来表示从坐标 (0, 0) 到坐标 (i, j)  的路径总数
    u(i, j) 表示坐标 (i, j) 是否可行，如果坐标 (i, j)(i,j) 有障碍物，u(i, j) = 1，否则 u(i, j) = 0。

    由于每次只能向下，或者向右走一步。
    因此从(0,0)到(i,j)的路径总和，取决于 (0,0)到(i-1,j)或(i,j-1)的路径总和

②状态转移：
    [动态转移方程]
        f(i,j) = 0                       u(i,j) = 1 有障碍
        f(i,j) = f(i-1,j) + f(i,j-1)     u(i,j) = 0 无障碍
    
   
③初始条件：
    由于开始位置可能也是障碍，所以要进行判断


④边界判断：
    如果障碍物在终点的那边堵住了，便不可能达到

⑤降维优化：
    [滚动数组思想] 是一种常见的动态规划优化方法
        当我们定义的状态在动态规划的转移方程中只和某几个状态相关的时候，就可以考虑这种优化方法，目的是给空间复杂度「降维」
        
        简单说就是压缩的维度是循环计算的历史。因为状态转移方程只跟前一行和前一列的相邻两个值有关，
        在循环嵌套里计算的时候，前一行的值就是上一个循环的计算结果，只需要考录列这一个维度就行了



        优化后的复杂度
            时间复杂度：O(mn)
            空间复杂度：O(n)。利用滚动数组优化，我们可以只用 O(n) 大小的空间来记录当前行的 f 值

*/