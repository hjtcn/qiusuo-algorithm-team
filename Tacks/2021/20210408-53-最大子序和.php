<?php
/*
 * @Descripttion:  周一的题，终于有空来看题 前几天不太舒服，不太想搞
 * @Author: tacks321@qq.com
 * @Date: 2021-04-02 18:03:45
 * @LastEditTime: 2021-04-08 17:55:30
 */

/*
 * @lc app=leetcode.cn id=53 lang=php
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (53.06%)
 * Likes:    3081
 * Dislikes: 0
 * Total Accepted:    465.7K
 * Total Submissions: 865.9K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1]
 * 输出：1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：nums = [-1]
 * 输出：-1
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：nums = [-100000]
 * 输出：-100000
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^5 
 * 
 * 
 * 
 * 
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 动态规划
    // 执行用时：24 ms, 在所有 PHP 提交中击败了40.85%的用户
    // 内存消耗：16.6 MB, 在所有 PHP 提交中击败了16.90%的用户
    function maxSubArray($nums) {
        // 获取数组长度
        $len = count($nums);
        
        // 遍历数组 
        for ($i=1; $i < $len; $i++) { 
            // 如果前一个值大于0 就加到当前值上
            if($nums[$i-1] > 0) {
                // 修改源数组
                $nums[$i] += $nums[$i-1];
            }
        }
        return max($nums);
    }
}

class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 动态规划
    // 执行用时：20 ms, 在所有 PHP 提交中击败了76.53%的用户
    // 内存消耗：16.4 MB, 在所有 PHP 提交中击败了87.79%的用户
    function maxSubArray($nums) {
        // 获取数组长度
        $len = count($nums);
        $max = $nums[0]; 
        
        // 遍历数组 
        for ($i=1; $i < $len; $i++) { 
            // 如果前一个值 > 0
            if($nums[$i-1] > 0) {
                // 修改源数组
                $nums[$i] += $nums[$i-1];
            }
            // 替换最大值
            if($nums[$i] > $max) {
                $max = $nums[$i];
            }
        }
        return $max;
    }
}
// @lc code=end



/*
【题目】
    最大子序和！ 说实话，这个题我校招之前见到过，还是比较经典的。

    找到一个具有最大和的连续子数组
    
【题解】
    动态规划
        假设 nums 数组的长度是 n，下标从 0 到 n-1。

        我们用 f(i) 代表以第 i 个数结尾的「连续子数组的最大和」
            f(i) = max { f(i-1) + num[i] }
 
        时间复杂度O(N)
        空间复杂度O(1)


*/