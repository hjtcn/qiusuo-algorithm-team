<?php
/*
 * @Descripttion: 周五快乐。可惜周末要加班，呜呜呜呜呜
 * @Author: tacks321@qq.com
 * @Date: 2021-03-05 14:52:00
 * @LastEditTime: 2021-03-05 14:56:53
 */

/*
 * @lc app=leetcode.cn id=704 lang=php
 *
 * [704] 二分查找
 *
 * https://leetcode-cn.com/problems/binary-search/description/
 *
 * algorithms
 * Easy (55.77%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    97.9K
 * Total Submissions: 175K
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的
 * target，如果目标值存在返回下标，否则返回 -1。
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 你可以假设 nums 中的所有元素是不重复的。
 * n 将在 [1, 10000]之间。
 * nums 的每个元素都将在 [-9999, 9999]之间。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function search($nums, $target) {
        $left = 0;
        $right= count($nums) - 1;
        // 经典二分
        while($left <= $right) {
            $mid = $left + floor( ($right - $left) / 2 );
            if($nums[$mid] < $target) {
                $left = $mid + 1;
            } elseif($nums[$mid] > $target) {
                $right= $mid - 1;
            } else {
                return $mid;
            }
        }
        // 不存在
        return -1;

    }
}
// @lc code=end

