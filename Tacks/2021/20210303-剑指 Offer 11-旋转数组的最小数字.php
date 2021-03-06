<?php

/*
 * @Descripttion: 过了周三 时间就快起来了
 * @Author: tacks321@qq.com
 * @Date: 2021-03-03 11:08:31
 * @LastEditTime: 2021-03-03 14:57:32
 */


/*
剑指 Offer 11. 旋转数组的最小数字

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。

例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0
注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

 */

 
class Solution {

    /**
     * @param Integer[] $numbers
     * @return Integer
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了91.89%的用户
    // 内存消耗：15.7 MB, 在所有 PHP 提交中击败了81.08%的用户
    function minArray($numbers) {
        $len  = count($numbers);
        $left = 0;
        $right= $len - 1;
        // 二分核心代码
        while($left < $right) {
            // 当然有个性质可以用一下 提前终止
            if ($numbers[$left] < $numbers[$right]) {
                return $numbers[$left];
            }

            $mid = $left + floor(($right-$left)/2); // 向上取整
            
            if($numbers[$mid] < $numbers[$right]) {
                $right = $mid;    
                // 细节点 （因为中间值mid是较小的，右边界可以直接用mid了，不能mid-1，这样有可能错过最小值）
                // 搜索区间变为 [left, mid]
            }elseif($numbers[$mid] > $numbers[$right]) {
                $left = $mid + 1; 
                // 细节点 （因为中间值mid已经比较了，比右边还大，肯定不是最小值，所以+1）
                // 搜索区间变为 [mid + 1, right]
            }else{
                $right--; 
                // 巧妙 （如果都是相等的，那么肯定不是最小值，我只需要缩小右边界即可）
            }
        }
        return $numbers[$left];
    }
}


/*
【题目】
    旋转数组：
        也就是说以某一个点阶段，将其移动到数组后面，而且数组是递增数组，那么也就是求源数组的第一个元素，即为最小

        那么旋转后的数组，最小值一定是第一个比开头元素小的元素

    特点：
        数组最后一个元素x
            最小值右侧的元素，它们的值一定都小于等于 x；
            最小值左侧的元素，它们的值一定都大于等于 x。

    遇到有序的数组，那么必然选择二分
【解法：二分】
    看了题解，一画图，就清晰了

    伪代码：
        left  = 0
        right = size(nums) - 1
        while(left <= right)
            mid   = left + (right-left)//2
            if nums[mid] < nums[right]
                说明，最小值在左半边
            elseif nums[mid] > nums[right]
                说明，最小值在右半边
            else nums[mid] == nums[right] 这个地方题解分析的很清楚
                等于的话，由于递增数组也有可能会有重复元素
                那么其实有点类似我们周一的哪个题《稀疏数组搜索》（如何处理空元素的过程一样）
                既然元素一样，那必然不是最小值，那我就缩小右边区间就好了  

【时间复杂度】
    平均时间复杂度为 O(logn)
    最差时间复杂度为 O(n)
    空间复杂度为 O(1)


哈哈哈哈哈哈哈哈哈哈，今天的题看了官方题解，觉得是讲的最清楚的一会，竟然看懂啦。
*/