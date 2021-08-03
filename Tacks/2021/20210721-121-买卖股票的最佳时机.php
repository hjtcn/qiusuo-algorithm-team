<?php
/*
 * @Descripttion: 周一的题
 * @Author: tacks321@qq.com
 * @Date: 2021-07-20 16:57:55
 * @LastEditTime: 2021-07-21 19:00:57
 */



/*
 * @lc app=leetcode.cn id=121 lang=php
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (56.84%)
 * Likes:    1722
 * Dislikes: 0
 * Total Accepted:    492.6K
 * Total Submissions: 863.4K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 * 
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 * 
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
    1 <= prices.length <= 10^5
    0 <= prices[i] <= 10^4
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        $maxprofit = 0; // 最大利润
        
        $size = count($prices);
        // 边界判断
        if($size <= 1) {
            return $maxprofit;
        }

        // 两层循环暴力处理
        for ($i=0; $i < $size-1; $i++) { 
            for ($j=$i+1; $j < $size; $j++) { 
                $curprofit = $prices[$j] - $prices[$i];
                // 最大利润更新
                if($curprofit > $maxprofit) {
                    $maxprofit = $curprofit;
                }
            }
        }
        return $maxprofit;

    }
}



class Solution2 {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    // 贪心：一次遍历
    // 执行用时：296 ms, 在所有 PHP 提交中击败了94.18%的用户
    // 内存消耗：27.7 MB, 在所有 PHP 提交中击败了17.46%的用户
    function maxProfit($prices) {
        $maxprofit = 0; // 最大利润
        $minprice  = 10000; // 当前最低值  0 <= prices[i] <= 10^4

        $size = count($prices);
        // 边界判断
        if($size <= 1) {
            return $maxprofit;
        }

        // 一次遍历
        for ($i=0; $i < $size; $i++) { 
            // buy 最低的买入
            $minprice = min($minprice, $prices[$i]); 
            // sell 最高的卖出
            $maxprofit = max($maxprofit, $prices[$i] - $minprice); 
            

            // if($prices[$i] < $minprice) {
            //     // buy 最低的买入
            //     $minprice = $prices[$i]; 
            // } else if($prices[$i] - $minprice > $maxprofit) {
            //     // sell 最高的卖出
            //     // 如果当前卖出的利润，大于利润，就进行卖出
            //     $maxprofit = $prices[$i] - $minprice;
            // }
        }
        return $maxprofit;

    }
}



class Solution3 {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    // 动规
    // 执行用时：364 ms, 在所有 PHP 提交中击败了64.20%的用户
    // 内存消耗：62 MB, 在所有 PHP 提交中击败了13.07%的用户
    function maxProfit($prices) {
        $size = count($prices);
        // 边界判断
        if($size <= 1) {
            return 0;
        }

        $dp[0][0] = 0;  // 开始,不持股
        $dp[0][1] = -$prices[0]; // 开始, 持股
        
        // 从第二天开始遍历
        for ($i=1; $i < $size ; $i++) { 
            // 前一天没有股票 ; 前一天有股票,今天卖出去   取大的
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][1] + $prices[$i]);
            // 前一天有股票 ; 前一天没有股票,今天买入     取大的
            $dp[$i][1] = max($dp[$i-1][1], -$prices[$i]);
        }
        return $dp[$size-1][0];
    }
}



// @lc code=end


$Obj = new Solution;
$Obj = new Solution2;
$Obj = new Solution3;
$prices = [7,1,5,3,6,4];
$res    = $Obj->maxProfit($prices);
var_dump($res);



/*

【 Clarification 】

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

题意：
    给一个数组 prices，数组中每个元素，表示第i天的股价，要求如何获得最大利润。
    利润：卖出-买入
    隐含条件：买入的价钱小于卖出，才能获得收益，没人是傻子明知道亏钱，还卖出。
              假设买入的为 prices[i] , 卖出的为 prices[j] （j>i）
思路：
    我们找到一个两者差距最大的即可，可以利用暴力两层循环处理

特别的:
    如果说  prices = [3,2,1]; 那么不发生交易,便是最好的结果
    
【 Possible-Solution 】
（1）暴力
        两层for循环，寻找 prices[j] , prices[i] 差值最大的时候进行返回
    复杂度：
        时间复杂度：O(n^2) 。 两层for循环 n(n-1)/2 次
        空间复杂度：O(1)
（2）贪心：一次遍历
        再仔细思考一下，为什么要两次for循环暴力，虽然比较好想，但是复杂度太高。
        我们主要是求差值最大，如果记作买入buy 卖出为sell
            如果我们希望 sell - buy 最大
            那么我们需要 buy最小，sell最大
        这个时候我们如果知道最低值buy，那么我们只需要在购入之后，有一个最高sell那不就行了
            所以可以采用一个变量 minprice 保存最低值，一次循环中不断更新这个变量，并且求解当前最大利润
    复杂度：
        时间复杂度：O(n) 
        空间复杂度：O(1)
（3）动态规划
        动规的思想，需要有一种能力，把问题的情况划分的特别清晰，归纳能力比较重要

        买卖股票有约束，根据题目意思，有以下两个约束条件：
            条件 1：你不能在买入股票前卖出股票；
            条件 2：最多只允许完成一笔交易。

        状态:
            dp[i][j] 下标为i这一天结束的时候,手上持股j时候,我们持有的现金数目
                j = 0 当前不持股
                j = 1 当前持股
            下标为 i 的这一天的计算结果包含了区间 [0, i] ，因此最后输出 dp[len - 1][0]。
        状态转移方程:
            dp[i][0]: 今天不持股
                1) 昨天不持股,今天什么也不做
                2) 昨天持股,今天卖出 获得收益 =======> 收钱
            dp[i][1]: 今天持股
                1) 昨天持股,今天什么也不错
                2) 昨天不持股,今天买入股票
        
            

     
*/