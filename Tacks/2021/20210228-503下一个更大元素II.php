<?php
/*
 * @lc app=leetcode.cn id=503 lang=php
 *
 * [503] 下一个更大元素 II
 */

 /*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
如果不存在，则输出 -1。

示例 1:

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数； 
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
  */
// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function nextGreaterElements($nums) {
        $arr = []; // 结果数组
        $size= count($nums); // 循环数组大小 
        for($i=0; $i<$size; $i++) {
            $j = $i + 1; // 下一个位置的元素
            $len = 0; 
            while(1) {
                // 如果超过了长度，就重置
                if($j >= $size) {
                    $j = 0;
                }
                // 后面的值如果大于当前值
                if($nums[$j] > $nums[$i]) {
                    $arr[] = $nums[$j];
                break;
                } else {
                    // 否则继续向后遍历
                    $j++;
                    $len++;
                }
                // 如果已经遍历过一轮，还找不到最大值
                if($len >= $size) {
                    $arr[] = -1;
                break;
                }
            }
        }
        return $arr;
    }
}
// @lc code=end

