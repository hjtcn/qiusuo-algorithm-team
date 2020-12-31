<?php

/*

题目: 面试题 17.10. 主要元素

数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。

示例 1：

输入：[1,2,5,9,5,9,5,5,5]
输出：5

说明：
你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-majority-element-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 执行用时：76 ms, 在所有 PHP 提交中击败了11.90%的用户
    // 内存消耗：20.2 MB, 在所有 PHP 提交中击败了41.67%的用户
    // 时间复杂度O(n) 
    function majorityElement1($nums) {
        $count = count($nums);
        // 边界判断
        if($count < 1) {
            return -1;
        }
        if($count == 1) {
            return $nums[0];
        }
        $map = []; // 用来存储统计过的值
        foreach($nums as $item) {
            if(isset($map[$item])) {
                $map[$item]++;
                if($map[$item] > $count/2) {
                    return $item;
                }
            }else{
                $map[$item] = 1;
            }
        }
        return -1;
    }

    // 摩尔投票法
    // 一对一抵消，核心思路就是如果有主要元素，那么最后留下来的那个一定是最多的
    // 执行用时：52 ms, 在所有 PHP 提交中击败了90.48%的用户
    // 内存消耗：20.2 MB 在所有 PHP 提交中击败了41.67%
    // 时间复杂度O(n)  空间复杂度O(1)
    function majorityElement2($nums) {
        $count = count($nums);
        // 边界判断
        if($count < 1) {
            return -1;
        }
        if($count == 1) {
            return $nums[0];
        }
        // 全程仅需要几个变量 不需要额外申请空间，因此空间复杂度O(1)
        $major = 0; // 假设当前是 主要元素
        $cur   = 0; // 当前票数
        foreach($nums as $item) {
            if($cur == 0) { // 如果当前票数为0
                $major = $item; // 那么重新指向一个值
            }
            // 核心算法
            if($major != $item) {
                // 如果主要元素 不是当前值  消除战术 有点像消消乐
                $cur--;
            }else{
                $cur++; // 如果是等于主要元素 就++
            }
        }
        // 如果最后票数为 正
        if($cur > 0) {
            // 统计一下是否超过一半的票数
            $number = 0;
            foreach($nums as $item) {
                if($item == $major) {
                    $number++;
                }
            }
            if($number > $count/2) {
                return $major;
            }
        }

        return -1;
    }
    
}