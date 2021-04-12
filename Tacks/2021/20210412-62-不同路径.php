<?php
/*
 * @Descripttion:  周一，冲冲冲
 * @Author: tacks321@qq.com
 * @Date: 2021-04-12 17:06:57
 * @LastEditTime: 2021-04-12 19:46:32
 */

/*
 * @lc app=leetcode.cn id=62 lang=php
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (63.94%)
 * Likes:    949
 * Dislikes: 0
 * Total Accepted:    234.9K
 * Total Submissions: 362.6K
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 * 
 * 问总共有多少条不同的路径？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：m = 3, n = 7
 * 输出：28
 * 
 * 示例 2：
 * 
 * 
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：m = 7, n = 3
 * 输出：28
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：m = 3, n = 3
 * 输出：6
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 题目数据保证答案小于等于 2 * 10^9
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    //【1】 动态规划
    // 执行用时：8 ms, 在所有 PHP 提交中击败了57.53%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了58.90%的用户
    function uniquePaths($m, $n) {
        // 声明二维数组 并且初始化
        $dp = [];
        // 左边界
        for ($i=0; $i < $m; $i++) { 
            $dp[$i][0] = 1;
        }
        // 上边界
        for ($j=0; $j < $n; $j++) { 
            $dp[0][$j] = 1;
        }

        // 遍历二维数组，计算每一格子
        for ($i=1; $i < $m; $i++) { 
            for ($j=1; $j < $n; $j++) { 
                $dp[$i][$j] = $dp[$i-1][$j] + $dp[$i][$j-1];
            }
        }

        // 返回最后一个值
        return $dp[$m-1][$n-1];
    }
}


class Solution2 {

    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    // 【2】优化
    // 执行用时：8 ms, 在所有 PHP 提交中击败了57.53%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了43.84%的用户
    function uniquePaths($m, $n) {
        $dp = [];
        // 初始化  纵向数列
        for ($i=0; $i < $n; $i++) { 
            $dp[$i] = 1;
        }
        // 遍历每一个位置
        for ($i=1; $i < $m; $i++) { 
            for ($j=1; $j < $n; $j++) { 
                // 将自身与上一格相加，得到新的右格子
                $dp[$j] = $dp[$j] + $dp[$j-1];
            }
        }
        return $dp[$n-1];
        
    }
}


 

// @lc code=end

/*
【题目】
    机器人每次只能向下或者向右移动一步。
    看看从 start 走到 finnish 可以有多少种走法！

直接看官方题解，跟着思路走！
【题解】
    每一格的路径，是由上一格和左一格决定的
        => 动态规划处理
    
    动态方程
        dp[i][j] = dp[i-1][j] + dp[i][j-1]

    特别的
        由于左边边界位置 第一行dp[0][j] 或者 第一列 dp[i][0] 所以路径可能性只能是1

    读到这里
        有点像之前数学中的归纳法
【伪代码】
    初始化矩阵
        建立 m*n 的矩阵，注意第0行和第0列的元素均为1
    遍历矩阵
        当前格子 dp[i][j]由其上边和左边的格子决定，这里对应动态方程
    遍历结束
        返回最后一个格子的数量 也就是 dp[m-1][n-1]

【复杂度】
    时间：O(m*n)  每一格子都要计算，双层循环
    空间：O(m*n)  申明了二维矩阵

【优化】
    只用长度为n的一维数组，来计算纵向列表
    
    每一次计算的时候，将自身与上一格相加，得到新的右格子

    空间：O(n)




    
【数学思维】
    组合数学 问题


*/