<?php
/*
 * @Descripttion: 周二奥里给
 * @Author: tacks321@qq.com
 * @Date: 2021-03-09 13:41:52
 * @LastEditTime: 2021-03-09 14:56:07
 */



/*
 * @lc app=leetcode.cn id=34 lang=php
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.11%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    217.3K
 * Total Submissions: 513.6K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 * 
 * 进阶：
 * 
 * 
 * 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * -10^9 
 * nums 是一个非递减数组
 * -10^9 
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：19.3 MB, 在所有 PHP 提交中击败了53.71%的用户
    function searchRange($nums, $target) {
        // 初始化范围
        $targetRange = [-1, -1];

        // 查看是否能找到第一个值
        $left = $this->_binarySearch($nums, $target);
      

        // 验证一下是否找到正确的 第一个值 leftIdx
        if($left == count($nums) || $nums[$left] != $target) {
            return $targetRange;
        }

        // 这里模拟找到第二个 （$target + 1）
        $right =  $this->_binarySearch($nums, $target + 1);

        // 最后第二个值 -1 
        $targetRange = [$left, $right - 1];
        
        // var_dump($targetRange);
        return $targetRange;
    }

    /**
     * 二分辅助方法
     *
     * @date  2021-03-09 14:06:29
     * @param [array] $nums
     * @param [int] $target
     */
    private function _binarySearch($nums, $target) {
        $low = 0;
        $high= count($nums) - 1;
        while($low <= $high) {
            $mid = ($low + $high) >> 1; // 左中值
            if($nums[$mid] < $target) {
                $low = $mid + 1;
            } elseif($nums[$mid] >= $target) {
                $high= $mid - 1;
            } 
        }
        //  这里找到左边的位置
        return $low;
        
    }
}
// @lc code=end


$obj = new Solution();

$nums  = [5,7,7,8,8,10];
$target= 8;

$nums  = [1];
$target= 1;
$obj->searchRange($nums, $target);

/*
【题目】
    给定一个按照升序排列的整数数组 nums，和一个目标值 target。
    

    找出给定目标值在数组中的开始位置和结束位置 =》 也就是这个目标值的也有可能连续重复出现
        也就是说，要找到这样两个值「第一个等于  target  的位置」「第一个大于 target 的位置减一」

    有序数组 + 时间复杂度为 O(log n)  =》 可以考虑一下二分
【题解】
    当然直接循环遍历也可以，但是复杂度是O(n) ,这里就不实现了。

    我们要找到的这两个值，分别记作 leftIdx 和 rightIdx

    我们首先利用二分的框架来找到第一个值 leftIdx
        如果找不到就可以直接返回[-1,-1]
    然后我们利用二分找到第二个值 rightIdx
        从而返回题的答案
    
    =》 抽离子方法（二分）
    
    关键是，我们如何用同一个二分代码，来找到两个位置的值，

    第二次的值，找到 $target + 1 的位置，然后再 最后计算的时候 -1


    但是感觉这个题的边界好难，淦

【时间复杂度】
    O(logn)

*/