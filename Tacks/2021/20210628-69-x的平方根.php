<?php
/*
 * @Descripttion: 周一的题，补一下 ，利用二分
 * @Author: tacks321@qq.com
 * @Date: 2021-06-29 16:57:51
 * @LastEditTime: 2021-06-29 17:32:10
 */

/*
 * @lc app=leetcode.cn id=69 lang=php
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (39.27%)
 * Likes:    703
 * Dislikes: 0
 * Total Accepted:    331.1K
 * Total Submissions: 843K
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
    // 二分
    // 执行用时：4 ms, 在所有 PHP 提交中击败了98.81%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了20.24%的用户
    function mySqrt($x) {
        if($x <= 0) {
            return 0;
        }
        $left = 0;
        $right= ceil($x/2); // 加速一下有边界
        while($left <= $right) {
            // $mid = $left + ceil(($right - $left) / 2);
            $mid = ($right + $left + 1) >> 1;
            $res = $mid*$mid;
            if($res < $x) {
                $left = $mid + 1;// 缩小左边界
            } elseif($res > $x) {
                $right = $mid - 1; // 缩小右边界
            } else {
                return $mid; // 刚好等于的情况 (mid*mid=x)
            }
        }
        // 如果遍历结束 当 right < left 的时候，我们返回right边界即可
        return $right;
    }
}
// @lc code=end

/*
【题目】
    就是找一个最大整数k  让 k*k <= x

【题解】
    二分思想
        条件：
            必须是一组有序的整数序列
        属性：        
            左边界：设置为0
            右边界：设置为最大值 x/2（但是我们可以知道 (x/2) * (x/2) <= x ）
            中间值：mid, mid = (right + left + 1) >> 1
        循环：
            二分每一步比较的时候都是缩小查找范围，通过比较  mid 与 x 的关系来调整上下边界范围。
        时间复杂度：
            可以将 O(n) 的时间复杂度降低到 O(log(n)), 算法思想很厉害， x就是查找的次数
        空间复杂度：
            O(1)
            
*/
