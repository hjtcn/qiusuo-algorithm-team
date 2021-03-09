<?php
/*
 * @Descripttion: 周一奥里给 女神节快乐
 * @Author: tacks321@qq.com
 * @Date: 2021-03-08 14:29:42
 * @LastEditTime: 2021-03-08 14:42:24
 */


/*
 * @lc app=leetcode.cn id=69 lang=php
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (39.09%)
 * Likes:    612
 * Dislikes: 0
 * Total Accepted:    267.2K
 * Total Submissions: 681.3K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了56.80%的用户
    // 内存消耗：14.9 MB, 在所有 PHP 提交中击败了100.00%的用户
    function mySqrt($x) {
        if($x <= 0) {
            return 0; // 边界判断
        }
        $left = 0;
        $right= $x;
        $ans  = 0;
        while($left <= $right) {
            $mid = ($left + $right + 1) >> 1;
            if($mid * $mid <= $x) {
                $ans = $mid;
                $left= $mid + 1; // 缩小左边界
            } else {
                $right= $mid - 1; // 缩小右边界
            }
        }
        return $ans;
    }
}
// @lc code=end

/*
【题目】
    考察二分查找

    实际开发中肯定用函数sqrt()
【二分】
    也就是要找到一个最大k来满足  k^2 <= x
    
    二分的下边界为0，上边界为x

    每次比较mid的平方跟x的关系，并在比较过程中，缩小边界范围

【复杂度】
    时间：O(logn)
    空间：O(1)

*/