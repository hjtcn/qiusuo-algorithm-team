<?php



/*
 * @Descripttion: 
 * @Author: tacks321@qq.com
 * @Date: 2021-04-09 16:24:20
 * @LastEditTime: 2021-04-09 16:49:35
 */


/*

[面试题 17.16. 按摩师]

一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

注意：本题相对原题稍作改动

 

示例 1：

输入： [1,2,3,1]
输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
示例 2：

输入： [2,7,9,3,1]
输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
示例 3：

输入： [2,1,4,5,3,1,1,3]
输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/the-masseuse-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */


class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 动态规划
    // 执行用时：4 ms, 在所有 PHP 提交中击败了85.71%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了78.57%的用户
    function massage($nums) {
        $len = count($nums);
        // 边界判断
        if($len == 0) {
            return 0;
        }
        // 考虑前 i 个预约，第 i 个预约不接的最长预约时间
        $dp0 = 0;
        // 表示考虑前 i 个预约，第 i 个预约接的最长预约时间
        $dp1 = $nums[0];

        for ($i=1; $i < $len; $i++) { 
            $tdp0 = max($dp0, $dp1); // 计算 dp[i][0]
            $tdp1 = $dp0 + $nums[$i]; // 计算 dp[i][1]

            // 变量更新
            $dp0 = $tdp0;
            $dp1 = $tdp1;
        }
        return max($dp0 , $dp1);
    }
}


 /*
【题目】
    给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
    不接待相邻的请求。

【解析】
    动态规划
    
    一共两种情况，到底接不接第i个人的预约，所以有 下面两种情况！！！

        dp[i][0] 表示考虑前 i 个预约，第 i 个预约不接的最长预约时间
        dp[i][1] 表示考虑前 i 个预约，第 i 个预约接的最长预约时间


        假设已知dp[i-1][0], dp[i-1][1]的值

        那么
        
        第i的预约不接
            由于这个状态下第 i 个预约是不接的，所以第 i-1 个预约接或不接都可以
            dp[i][0] = max(dp[i-1][0], dp[i-1][1])

        第i的预约接
            对于dp[i][1] 这个状态 i 个预约要接，但是相邻不能接
            dp[i][1] = dp[i-1][0] + nums[j]
        
        最终
            max(dp[n][0], dp[n][1]) n表示预约个数


    所以我们发现我们想要知道 dp[i][0] 或者 dp[i][1] 都首先要知道 dp[i-1][0] dp[i-1][1]

    所以我们可以找几个变量来声明，而不用数组。

    
    时间复杂度 O(n)
    空间复杂度 O(1)


 */