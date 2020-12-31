<?php

/*
 * @lc app=leetcode.cn id=503 lang=php
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode-cn.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (57.62%)
 * Likes:    212
 * Dislikes: 0
 * Total Accepted:    33.6K
 * Total Submissions: 58K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x
 * 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数； 
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 * 
 * 
 * 注意: 输入数组的长度不会超过 10000。
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    // 执行用时：1472 ms, 在所有 PHP 提交中击败了16.67%的用户
    // 内存消耗：17.8 MB, 在所有 PHP 提交中击败了28.57%的用户
    function nextGreaterElements1($nums) {
        $arr = [];// 结果数组
        $size= count($nums);
        for($i=0; $i<$size; $i++) {
            $j = $i + 1; // 下一个元素的位置
            $len = 0;
            while(1) {
                // 如果超过长度，就重置
                if($j >= $size) {
                    $j = 0;
                }
                // 后面比当前值大
                if($nums[$j] > $nums[$i]) {
                    $arr[] = $nums[$j];
                    break;
                }else{
                    // 否则继续向后遍历
                    $j++;
                    $len++;
                }
                // 如果已经遍历过一轮，找不到大于的值
                if($len >= $size) {
                    $arr[] = -1;
                    break;
                }
            }

        }
        return $arr;

    }

    // 单调栈
    public function nextGreaterElements($nums)
    {

    }

}
// @lc code=end


/*

【1】采用双指针遍历法

控制好终止条件
    i标识当前元素，j标识下一个更大的元素的位置。
        当j遍历到数组末尾，将重新回执到0位置
        当j的位置元素，大于i的元素，直接存储到结果数组中，跳出循环
        如果j已经遍历过一轮，也就是整个数组的长度，那么就直接标记为-1
可见这样是非常慢的，因为你每次都要遍历整个数组来判断有没有最大值，太费时间。


【2】单调栈

觉得这个单调栈出现频率挺高了 有点不太好搞


*/