<?php
/*
 * @Descripttion: 周四来刷个题
 * @Author: tacks321@qq.com
 * @Date: 2021-06-24 17:42:53
 * @LastEditTime: 2021-06-24 18:39:32
 */


/*
 * @lc app=leetcode.cn id=198 lang=php
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Medium (49.86%)
 * Likes:    1516
 * Dislikes: 0
 * Total Accepted:    316.8K
 * Total Submissions: 633.4K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
 * 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
 * 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 * 
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 * 
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
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了73.01%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了80.37%的用户
    function rob($nums) {
        // 边界判断
        if(!$nums) {
            return 0;
        }
        $count = count($nums);
        if($count == 1) {
            return $nums[0];
        }
        if($count == 2) {
            return max($nums[0], $nums[1]);
        }

        // 初始化dp数组
        $dp = array_fill(0, $count, 0); // 全部填充0
        $dp[0] = $nums[0];
        $dp[1] = max($nums[0],$nums[1]);
        
        // 遍历一下数组
        for ($i=2; $i < $count; $i++) { 
            // 动态转移方程
            $dp[$i] = max($dp[$i-2] + $nums[$i], $dp[$i-1]);
        }
        
        return $dp[$count-1];
    }
}

class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了73.01%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了82.82%的用户
    function rob($nums) {
        // 边界判断
        if(!$nums) {
            return 0;
        }
        $count = count($nums);
        if($count == 1) {
            return $nums[0];
        }
        if($count == 2) {
            return max($nums[0], $nums[1]);
        }

        // 初始化 dp 变量
        $first = $nums[0];
        $second= max($nums[0], $nums[1]);

        // 遍历一下数组
        for ($i=2; $i < $count; $i++) { 
            $tmp = $second; 
            // 动态转移方程
            $second = max($first+$nums[$i], $second);
            $first  = $tmp;
        }
        
        return $second;
    }
}
// @lc code=end

/*
【题目】
    就是不能偷相邻两个房屋，每个房屋都有非负金额，看看可以最多可以偷多少。
    
【解析】
    动态规划 万物皆可DP

思考
1) 如果只有一间房，那么偷窃该房屋，即可获得最大金额。
2) 如果有两间房，那么直接比较两个房屋的金额，选择其中一间偷窃即可。
3) 如果有大于两间房以上，如何可以偷窃到最大金额？
    如果房屋数量为 k (k>2)
        两种情况。要么选择第k间房子，要么不选择
        a、偷窃第k间房子，就不能偷窃第k-1房屋。那么偷窃总金额= 前k-2间房屋的最高总金额 + 第k间房屋的金额之和。
        b、不偷第k间房子，那么偷窃总金额 = 前k-1间房屋的最高总金额。
    再两种方案中选择最大的选项，则为 k 间 房屋偷窃最大金额。
        利用dp[i]表示前i-1间房屋可以偷窃的最高总金额
    动态转移方程：
        dp[0] = nums[0]                         当i=0的时候，一间房
        dp[1] = max(nums[0], nums[1])           当i=1的时候, 两间房
        dp[i] = max(dp[i-2] + nums[i], dp[i-1]) i-1间房屋

    优化
        每间房屋的最高总金额只和该房屋的前两间房屋的最高总金额相关，
        因此可以使用滚动数组，在每个时刻只需要存储前两间房屋的最高总金额。
        i 其实只依赖于 i - 1 和 i - 2 两个状态

        first  保存 dp[i-2]
        second 保存 dp[i-1]

        然后在每一次的遍历中进行赋值替换。

    复杂度
        时间复杂度：O(n)，其中 n 是数组长度。只需要对数组遍历一次。
        空间复杂度：O(1)。优化后，只需要几个变量，即可。

*/