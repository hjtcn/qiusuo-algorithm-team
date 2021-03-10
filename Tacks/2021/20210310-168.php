<?php
/*
 * @Descripttion: 周三 脖子不太舒服
 * @Author: tacks321@qq.com
 * @Date: 2021-03-10 17:54:13
 * @LastEditTime: 2021-03-10 18:20:36
 */


/*
 * @lc app=leetcode.cn id=162 lang=php
 *
 * [162] 寻找峰值
 *
 * https://leetcode-cn.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (48.17%)
 * Likes:    390
 * Dislikes: 0
 * Total Accepted:    79.6K
 * Total Submissions: 163.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 峰值元素是指其值大于左右相邻值的元素。
 * 
 * 给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
 * 
 * 你可以假设 nums[-1] = nums[n] = -∞ 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,1]
 * 输出：2
 * 解释：3 是峰值元素，你的函数应该返回其索引 2。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,1,3,5,6,4]
 * 输出：1 或 5 
 * 解释：你的函数可以返回索引 1，其峰值元素为 2；
 * 或者返回索引 5， 其峰值元素为 6。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -2^31 
 * 对于所有有效的 i 都有 nums[i] != nums[i + 1]
 * 
 * 
 * 
 * 
 * 进阶：你可以实现时间复杂度为 O(logN) 的解决方案吗？
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了42.86%的用户
    // 内存消耗：15.5 MB, 在所有 PHP 提交中击败了7.14%的用户
    function findPeakElement($nums) {
        $left = 0;
        $right= count($nums) - 1;

        // 二分空壳
        while($left < $right) {
            $mid = ($left  + $right ) >> 1; // 左中值
            // 核心思路
            if($nums[$mid] > $nums[$mid+1]){
                // 降序过程
                // 过滤右边一部分
                $right = $mid;
            } else {
                // 升序过程 $nums[$mid] <= $nums[$mid+1]
                // 过滤左边一部分
                $left  = $mid + 1;
            }
            // 当 left = right
            // 就找到了
        }
        return $left;
    }
}
// @lc code=end

/*
【题目】
    峰值元素是指其值大于左右相邻值的元素。         =》 部分有序
    对于所有有效的 i 都有 nums[i] != nums[i + 1]   =》 保证不会有重复的

        如果是升序过程  第一个元素小于第二个元素 那么不会峰值
        如果是降序过程  第一个元素大于第二个元素 那么也不会有峰值
        如果有峰值  那么出现在中间某个位置 达到峰值 那么当前会有位置元素 大于下一个元素值
    
【题解】
    1、遍历思路
        O(n)
        那么当前位置如果是峰值，那么就是 当前值 大于 后面的值 （nums[n] > nums[n+1]）

        通过这样思路即可解题
    2、二分思路
        O(logn)
        既然数组有峰值，那么就是部分有序

        我们就利用mid值来定位峰值的位置
            其实思路还是（nums[n] > nums[n+1]）
        只不过将线性遍历，转化为跳步来找，就是二分每次会过滤一半的元素，从而节省时间


        // 数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

        当然这个题目说可以有多个峰值，所以我们只要找到一个就行，而且是肯定有峰值

        有几种情况
        ① 一个峰值
            nums = [1,2,4,1] 
            第一次二分
                mid = 1
                nums[1] < nums[2] 2<4  上升阶段，过滤掉左边部分
                left = mid + 1  = 2
            第二次二分
                mid = 2
                nums[2] > nums[3] 2>1  下降阶段，过滤掉右边部分
                right = mid = 2
            此时left == right = 2 二分结束
            返回 4
        ② 多个峰值
            nums = [1,2,3,1,2,4,5,1] 
            第一次二分
                mid = 3
                nums[3] < nums[4] 1<2 上升阶段， 过滤掉左边部分，也就是[1,2,3,1] 这一块
                那么 这个时候就过滤掉左边部分的峰值，但是没关系，还有右边的
                left = mid + 1 = 4
            第二次二分
                mid = 5
                nums[5] < nums[6] 4<5 上升阶段，过滤掉做左边部分，也就是[2,4] 这一块
                left = mid + 1 = 6
            第三次二分
                就找到了 
                

    
*/