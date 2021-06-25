<?php
/*
 * @Descripttion: 终于周五咯
 * @Author: tacks321@qq.com
 * @Date: 2021-06-25 15:49:02
 * @LastEditTime: 2021-06-25 17:42:14
 */

/*
 * @lc app=leetcode.cn id=213 lang=php
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (42.70%)
 * Likes:    712
 * Dislikes: 0
 * Total Accepted:    130.6K
 * Total Submissions: 305.3K
 * Testcase Example:  '[2,3,2]'
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈
 * ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
 * 
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,2]
 * 输出：3
 * 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *  
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3,1]
 * 输出：4
 * 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：0
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
    // 动态规划
    // 执行用时：4 ms, 在所有 PHP 提交中击败了94.62%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了90.32%的用户
    function rob($nums) {
        // 边界判断
        if(empty($nums)) {
            return 0;
        }
        $n = count($nums);
        if(1 == $n) {
            return $nums[0];
        }
        if(2 == $n) {
            return max($nums[0], $nums[1]);
        }

        // 由于是环状的
        // 判断选择第一间房子
        // 或者选择最后一间房子 哪个的dp总金额值更大
        return max($this->_helper( array_slice($nums, 0, $n-1) ), $this->_helper( array_slice($nums, 1, $n) ) );

    }

    /**
     * 辅助函数：动规
     *
     * @date  2021-06-25 16:56:18
     * @param array $nums
     * @return int
     */
    private function _helper($nums) {
        $count = count($nums);
        // dp的大小
        $dp = array_fill(0, $count, 0); // 全部填充0
        $dp[0] = $nums[0];
        $dp[1] = max($nums[0],$nums[1]);
        
        // 遍历一下数组
        for ($i=2; $i < $count; $i++) { 
            // 动态转移方程
            $dp[$i] = max($dp[$i-2] + $nums[$i], $dp[$i-1]);
        }
        return $dp[$count - 1];

    }

}





class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 动态规划
    function rob($nums) {
        // 边界判断
        if(empty($nums)) {
            return 0;
        }
        $n = count($nums);
        if(1 == $n) {
            return $nums[0];
        }
        if(2 == $n) {
            return max($nums[0], $nums[1]);
        }

        // 由于是环状的
        // 判断选择第一间房子
        // 或者选择最后一间房子 哪个的dp总金额值更大
        return max($this->_helper( array_slice($nums, 0, $n-1) ), $this->_helper( array_slice($nums, 1, $n) ) );

    }

    /**
     * 辅助函数：动规
     *
     * @date  2021-06-25 16:56:18
     * @param array $nums
     * @return int
     */
    private function _helper($nums) {
        $count = count($nums);

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




$Obj = new Solution();

$nums = [2,3,2];
$nums = [1,2,1,1];
echo $Obj->rob($nums);


// @lc code=end

/*
【题目】

    与昨天题是同系列题目：[198] 打家劫舍  https://leetcode-cn.com/problems/house-robber/description/

    
    Ⅰ：就是不能偷相邻两个房屋，每个房屋都有非负金额，看看可以最多可以偷多少。
    Ⅱ：首尾房屋相连，第一间和最后一间不能同时偷。
    
    
【解析】
    动态规划 万物皆可DP

思考

① 边界问题
    ⅠⅡ的处理办法是一样的，只能偷一间。
    1) 如果只有一间房，那么偷窃该房屋，即可获得最大金额。
    2) 如果有两间房，那么直接比较两个房屋的金额，选择其中一间偷窃即可。

② 动态转移方程
    
    如果有大于两间房以上，如何可以偷窃到最大金额？
        首尾相连问题需要考虑：第一间与最后一间不能同时偷！
            1）如果偷窃了第一间房屋，则不能偷窃最后一间房屋，因此偷窃房屋的范围是第一间房屋到最后第二间房屋；
            2）如果偷窃了最后一间房屋，则不能偷窃第一间房屋，因此偷窃房屋的范围是第二间房屋到最后一间房屋。

    假设 nums数组长度为 n
        如果偷第一间房屋，那么偷窃范围在 [0, n-2]
        如果偷最后一间房屋，那偷窃范围在 [1, n-1]
    假设偷窃范围为 [start, end]。
    那么用dp[i] 表示从 [start,i]内可以偷窃的最高总金额
    于是
        dp[i] = max(dp[i-2] + nums[i], dp[i-1])     （ 是不是与昨天动规转移方程一样）
    边界条件
        dp[0] = nums[0]                         当i=0的时候，一间房
        dp[1] = max(nums[0], nums[1])           当i=1的时候, 两间房

    那我们比较一下
        [start,end]
            当 [0, n-2] 
            当 [1, n-1]
        这两种情况得到的 dp大小。

        核心代码都没有变化
③ 降维

 
    复杂度
        时间复杂度：O(n)，其中 n 是数组长度。只需要对数组遍历两次。
        空间复杂度：O(1)。优化后，只需要几个变量，即可。


*/