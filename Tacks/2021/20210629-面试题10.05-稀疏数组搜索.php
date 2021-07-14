<?php
/*
 * @Descripttion: 周二 老题目了
 * @Author: tacks321@qq.com
 * @Date: 2021-06-29 18:01:41
 * @LastEditTime: 2021-06-29 18:16:07
 */


/*
面试题 10.05. 稀疏数组搜索

稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。
示例2:

 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4
提示:

words的长度在[1, 1000000]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sparse-array-search-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 
 */


class Solution {

    /**
     * @param String[] $words
     * @param String $s
     * @return Integer
     */
    // 二分
    // 执行用时：20 ms, 在所有 PHP 提交中击败了80.00%的用户
    // 内存消耗：16.4 MB, 在所有 PHP 提交中击败了40.00%的用户
    function findString($words, $s) {
        // 边界
        $left = 0;
        $right= count($words) - 1;
        
        while($left <= $right) {
            $mid = $left + floor(($right - $left) / 2); // 中间位置
            
            // 解决空格的问题
            if($words[$mid] == "") {
                // 如果边界等于 s
                if($words[$right] == $s) {
                    return $right;
                }else{
                    // 如果边界不等于 s ,向左缩小范围
                    $right--;
                }


            // 标准二分
            }elseif($words[$mid] == $s){
                // 整好等于中间位置的值
                return $mid;
            }elseif($words[$mid] > $s) {
                // 向左移动
                $right = $mid-1;
            }else{
                // 向右移动
                $left  = $mid+1;
            }
        }
        return -1;
    }
}

/*

【题目】
题意：
    在一个有序的字符串数组中，给一个字符串，定位出字符串所在的下标的位置
关键：
    字符串排好序，也就是说 a,b,c 这样的字典序, 字符串比较按照第一个字符进行依次比较。

【解析】
1、暴力
    直接遍历字符串数组，依次比较每一个值，是否等于目标值，如果是便返回下标

2、二分（有序字符串数组）
    核心：
        有序的数据结构（如数组，可以随机访问元素，确定范围），根据每次的折半，而达到缩减规模的效果。
    关键：
        边界的判断，多一个等号少一个都不一样，一定要注意。
    中间值的算法：
        后两种都可以
        mid = floor((left+right) / 2)
        对于PHP弱语言来说，int类型超过范围会转化成float，为了避免left+right大数相加
        mid = left + floor((right - left) / 2)
        mid = (right + left + 1) >> 1
    时间复杂度：
        通常使用二分后可以将O(n)的时间复杂度转化为O(logn)

    难点：
        主要是因为多了很多空字符串，导致不知道如何判断左移还是右移动

        这里，如果mid==""，那么我们判断一下right右边界的值，是否等于目标值s
            如果等于可以直接返回
            如果不等于，那么我们可以将right--,从而达到缩小范围的目的。
        剩下就是标准二分啦
*/