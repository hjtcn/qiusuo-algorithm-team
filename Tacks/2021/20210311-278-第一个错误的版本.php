<?php

/*
 * @Descripttion: 周四，冲冲冲，今天晚上 去看脆莓
 * @Author: tacks321@qq.com
 * @Date: 2021-03-11 14:00:10
 * @LastEditTime: 2021-03-11 14:11:26
 */

/*
 * @lc app=leetcode.cn id=278 lang=php
 *
 * [278] 第一个错误的版本
 *
 * https://leetcode-cn.com/problems/first-bad-version/description/
 *
 * algorithms
 * Easy (41.76%)
 * Likes:    262
 * Dislikes: 0
 * Total Accepted:    83K
 * Total Submissions: 195.5K
 * Testcase Example:  '5\n4'
 *
 * 
 * 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
 * 
 * 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
 * 
 * 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version
 * 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
 * 
 * 示例:
 * 
 * 给定 n = 5，并且 version = 4 是第一个错误的版本。
 * 
 * 调用 isBadVersion(3) -> false
 * 调用 isBadVersion(5) -> true
 * 调用 isBadVersion(4) -> true
 * 
 * 所以，4 是第一个错误的版本。 
 * 
 */

// @lc code=start
/* The isBadVersion API is defined in the parent class VersionControl.
      public function isBadVersion($version){} */

class Solution extends VersionControl {
    /**
     * @param Integer $n
     * @return Integer
     */
    // 执行用时：56 ms, 在所有 PHP 提交中击败了87.50%的用户
    // 内存消耗：14.9 MB, 在所有 PHP 提交中击败了100.00%的用户
    function firstBadVersion($n) {
        $left = 1;
        $right= $n;
        // 二分框架
        while($left < $right) {
            $mid = ($left+ $right) >> 1;
            // 调用方法 判断是否正常版本
            if($this->isBadVersion($mid)) {
                $right = $mid;
            }else{
                $left = $mid + 1;
            }
        }
        return $left;
    }
}
// @lc code=end

/*
【题目】
    假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

    其实就是在二分的基础上，增加了场景应用题，通过调用isBadVersion()来进行判断

【题解】
    也就是说，如果这个存在，比如 badIdx
    那么 i < badIdx  全部是正常的
         i >= badIdx 全部是错误的
    
    二分
        如果 nums[mid] error 错误版本 要缩小右边界，从而找到第一个错误的值
            right = mid         
        如果 nums[mid] yes  正常版本
            left = mid + 1
        

*/