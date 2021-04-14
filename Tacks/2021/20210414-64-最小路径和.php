<?php
/*
 * @Descripttion: 赶紧补一下周二的题  现在越来越严格了
 * @Author: tacks321@qq.com
 * @Date: 2021-04-14 10:36:21
 * @LastEditTime: 2021-04-14 11:40:58
 */


/*
 * @lc app=leetcode.cn id=64 lang=php
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (68.00%)
 * Likes:    847
 * Dislikes: 0
 * Total Accepted:    203.6K
 * Total Submissions: 297.9K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 * 
 * 说明：每次只能向下或者向右移动一步。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
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
 * 0 
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    // 执行用时：36 ms, 在所有 PHP 提交中击败了56.72%的用户
    // 内存消耗：18.7 MB, 在所有 PHP 提交中击败了32.84%的用户
    function minPathSum($grid) {
        // 边界判断
        if(!$grid || (isset($grid[0]) && !$grid[0])) {
            return 0;
        }

        // 获取网格大小
        $m = count($grid);
        $n = count($grid[0]);

        // 声明二维数组 并且初始化
        $dp = [];

        $dp[0][0] = $grid[0][0];

        // 左边界
        for ($i=1; $i < $m; $i++) { 
            $dp[$i][0] = $dp[$i-1][0] + $grid[$i][0];
        }
        // 上边界
        for ($j=1; $j < $n; $j++) { 
            $dp[0][$j] = $dp[0][$j-1] + $grid[0][$j];
        }

        // 遍历二维数组，计算每一格子
        for ($i=1; $i < $m; $i++) { 
            for ($j=1; $j < $n; $j++) { 
                // 状态转移方程
                // 当前格子的值+前一位的最小值
                $dp[$i][$j] = $grid[$i][$j] + min($dp[$i-1][$j], $dp[$i][$j-1]);
            }
        }

        // 返回最后一个值
        return $dp[$m-1][$n-1];
    }
}
// @lc code=end

/*
【题目】
    又是 m*n 的网格，跟昨天的题有点类似，但是不一样的是，这次每个网格都有路径值

    然后只能向右或者向下移动一步。

    最终求路径之和做小的情况。

【解析】
    最小路径：可以想到 两种做法，搜索（复杂度高）、动规。
    
    这里采用动态规划
        状态
            dp[i][j] 表示从 dp[i][j] 走到右下角的 (m-1,n-1) 的路径和。
        状态转移方程
            由于只能向右或者向下走 
                所以 dp[i][j] 下一步的状态 为dp[i+1][j]  或者 dp[i][j+1]
                那么我们如果让我们所有路径值最小 =》 就是让上一步的路径最小  （有点动态规划意思）
                也就是 min(dp[i+1][j], dp[i][j+1])
            
            dp(i,j) = grid(i,j) + min{ dp(i+1, j), dp(i,j+1) } 
    
    具体分析
        那么还是 m*n 的网格，必然就有边界，有边界就只能向一个方向移动，所以是固定的

        到达右下角的最小路径，要看前一个格子的最小路径。

        所以基本改改昨天的代码就可以实现。

【复杂度】
    时间：O(m*n)
    空间：O(m*n)
*/