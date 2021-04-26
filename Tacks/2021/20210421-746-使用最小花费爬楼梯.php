<?php
/*
 * @Descripttion: 周二
 * @Author: tacks321@qq.com
 * @Date: 2021-04-21 09:27:11
 * @LastEditTime: 2021-04-21 10:40:49
 */


/*
 * @lc app=leetcode.cn id=746 lang=php
 *
 * [746] 使用最小花费爬楼梯
 *
 * https://leetcode-cn.com/problems/min-cost-climbing-stairs/description/
 *
 * algorithms
 * Easy (54.58%)
 * Likes:    536
 * Dislikes: 0
 * Total Accepted:    91.3K
 * Total Submissions: 166.4K
 * Testcase Example:  '[10,15,20]'
 *
 * 数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
 * 
 * 每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
 * 
 * 请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：cost = [10, 15, 20]
 * 输出：15
 * 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
 * 输出：6
 * 解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * cost 的长度范围是 [2, 1000]。
 * cost[i] 将会是一个整型数据，范围为 [0, 999] 。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[] $cost
     * @return Integer
     */
    // 动态规划
    // 执行用时：12 ms, 在所有 PHP 提交中击败了96.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了52.00%的用户
    function minCostClimbingStairs($cost) {
        $len = count($cost);  // 看一下有几层楼
        
        // 初始化
        $dp = array_fill(0, $len + 1, false);
        
        $dp[0] = 0;
        $dp[1] = 0;

        for ($i=2; $i <= $len; $i++) { 
            // 动态转移方程
            $dp[$i] = min($dp[$i-1] + $cost[$i-1], $dp[$i-2] + $cost[$i-2]);
        }
        return $dp[$len];

    }
}

class Solution2 {

    /**
     * @param Integer[] $cost
     * @return Integer
     */
    // 优化空间
    // 执行用时：16 ms, 在所有 PHP 提交中击败了72.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了80.00%的用户
    function minCostClimbingStairs($cost) {
        $len = count($cost);  // 看一下有几层楼
        
        $prev = 0;// 到达第1级的cost
        $curr = 0;// 到达第2级cost
        $next = 0;
        // 到达第3级到第n级（最后一级）
        for ($i=2; $i <= $len; $i++) { 
            $next = min($curr + $cost[$i-1] , $prev + $cost[$i-2]);
            // 更新前两个值
            $prev = $curr;
            $curr = $next;
        }
        return $curr;
    }
}
// @lc code=end

/*

【题目】

    假设数组cost的长度n,  n 个阶梯分别对应下标 0 到 n-1, 楼层顶部对应下标 n , 问题等价于计算达到下标 n 的最小花费。

【解析】
    动态规划
        创建长度为n+1的数组dp，其中dp[i]表示达到下标i的最小花费。
            当i=0 或者 i=1  
                由于可以选择下标 0 或 1 作为初始阶梯 => dp[0]=0   dp[1]=0
            当 2<= i <= n 
                可以选择向上爬一个阶梯或者爬两个阶梯 
                => 那么可以两种方式  
                    1、使用下标 i-1 使用 dp[i-1]的花费到达 i
                    2、使用下标 i-2 使用 dp[i-2]的花费到达 i
                为了让总花费最小
                    应当取上面两种情况的最小值
                    dp[i] = min(dp[i-1] + cost[i-1] ,   dp[i-2] + cost[i-2])
        这样就得到动态转移方程


    时间复杂度：O(n)
    空间复杂度：O(n)
【优化】
    通常这样我们动态规划都会根据对应的数据结构，创建对应的 一维、二维、等数组
    
    但是一般我们可以将其转化
        例如 一维数组 => 多个变量实现
             二维数组 => 一维数组实现
        从而让空间复杂度达到一定程度的降维
    
    比如这个题：
    我们发现 i>=2 的时候 dp[i] 只跟前两个值有关系
    可以选择向上爬一个阶梯或者爬两个阶梯 
            dp[i] = dp[i-1] &  dp[i-2] 
    所以我们可以使用几个变量来表示
            next = curr & prev

    时间复杂度O(n)
    空间复杂度O(1)

【总结】
    动态规划基本流程
        1、发现子问题
        2、写出动态转移方程
        3、伪代码描述解题过程
        4、编码实现
        5、降维优化数据结构（可选）
    
【发散】
    题目中cost的数组，可以直接利用起来，通过累加的方式来计算到达楼顶的花费。（巧用题中的变量，真是省空间）

*/