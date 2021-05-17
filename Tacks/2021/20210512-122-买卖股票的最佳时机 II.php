<?php
/*
 * @Descripttion: 周一的题
 * @Author: tacks321@qq.com
 * @Date: 2021-05-12 14:31:45
 * @LastEditTime: 2021-05-12 16:03:38
 */


/*
 * @lc app=leetcode.cn id=122 lang=php
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (66.11%)
 * Likes:    1216
 * Dislikes: 0
 * Total Accepted:    361.8K
 * Total Submissions: 535.3K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 * 
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: prices = [7,1,5,3,6,4]
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: prices = [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: prices = [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    // 贪心
    // 执行用时：24 ms, 在所有 PHP 提交中击败了31.40%的用户
    // 内存消耗：16.8 MB, 在所有 PHP 提交中击败了95.90%的用户
    function maxProfit($prices) {
        $ans = 0;
        $size = count($prices);

        // 最后一天肯定要卖出去
        for ($i=1; $i < $size; $i++) { 
            # 只要后一天比前一天大，也就是正数，就可以搞
            $ans += max(0, $prices[$i] - $prices[$i-1] );
            // 但是要注意，这个操作并不是代表交易的次数
        }
        return $ans;
    }
}


class Solution2 {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    // 动规
    // 执行用时：36 ms, 在所有 PHP 提交中击败了7.85%的用户
    // 内存消耗：26.1 MB, 在所有 PHP 提交中击败了6.48%的用户
    function maxProfit($prices) {
        $size = count($prices);
        $dp[0][0] = 0;
        $dp[0][1] = 0-$prices[0];

        for ($i=1; $i < $size; $i++) { 
            # 动态转移方程
            // 前一天已经没有股票，前一天卖出股票大
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][1] + $prices[$i]);
            // 前一天还有股票，前一天已经没有股票
            $dp[$i][1] = max($dp[$i-1][1], $dp[$i-1][0] - $prices[$i]);

        }
        return $dp[$size-1][0];
    }
}

class Solution3 {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    // 动规 优化空间
    // 执行用时：20 ms, 在所有 PHP 提交中击败了65.87%的用户
    // 内存消耗：17 MB 在所有 PHP 提交中击败了77.82%的用户
    function maxProfit($prices) {
        $size = count($prices);
        $dp0 = 0;
        $dp1 = 0 - $prices[0];
        for ($i=1; $i < $size; $i++) { 
            // 只需要将 dp[i−1][0] 和 dp[i−1][1] 存放在两个变量中，
            $newdp0 = max($dp0, $dp1 + $prices[$i]);
            $newdp1 = max($dp1, $dp0 - $prices[$i]);
            $dp0 = $newdp0;
            $dp1 = $newdp1;
        }
        return $dp0;
    }
}
// @lc code=end




/*
【题目】
    尽可能地完成更多的交易（多次买卖一支股票），从而得到利益的最大化
        必须在再次购买前出售掉之前的股票，也就是只能有一股

    实际上就是找这个 数组，不相交的差值最大化
        而且交易没有次数限制，反正只要后面的数大于前面，就搞！

【解析】
    1、贪心
        prices 数组
        Li 表示 在第 i 天 买入
        Ri 表示 在第 i 天 卖出
        那么所得利润 prices[Ri] - prices[Li]

        贪心 ，可以简化为找 x 个长度为 1 的区间使得  =>   prices[Li+1] - prices[Li] 价值最大化
        需要说明的是，贪心算法只能用于计算最大利润，计算的过程并不是实际的交易过程。


        时间复杂度：O(n)，其中 n 为数组的长度。我们只需要遍历一次数组即可。
        空间复杂度：O(1)。只需要常数空间存放若干变量。


    2、动态规划
        看小马姐说万物皆可动规，那我也搞一下试试
        

        定义状态 
            dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润
            dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润

            dp[i][0] 的转移方程，
                    如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即 dp[i−1][0]
                    如果前一天结束的时候手里持有一支股票，dp[i−1][1]，这时候我们要将其卖出，并获得 prices[i] 的收益
                因此为了收益最大化，我们列出如下的转移方程
                    dp[i][0] = max{dp[i−1][0], dp[i−1][1] + prices[i]}
            dp[i][1] 的转移方程，
                    可能的转移状态为前一天已经持有一支股票，即 dp[i−1][1]，
                    或者前一天结束时还没有股票，即 dp[i−1][0]，
                这时候我们要将其买入，并减少 prices[i] 的收益。可以列出如下的转移方程：
                    dp[i][1]= max{dp[i−1][1], dp[i−1][0]−prices[i]}
            
        据状态定义我们可以知道第 0 天交易结束的时候 dp[0][0]=0, dp[0][1] = −prices[0]。


            持有股票的收益一定低于不持有股票的收益，因此这时候 dp[n−1][0] 的收益必然是大于p[n−1][1] 的，最后的答案即为dp[n−1][0]

【总结】
    贪心贪心，就是找到一个从当前来看是最优的解法
 
    
            
 

*/