/*
 * @lc app=leetcode.cn id=628 lang=php
 *
 * [628] 三个数的最大乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (50.42%)
 * Likes:    171
 * Dislikes: 0
 * Total Accepted:    27.3K
 * Total Submissions: 54.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [1,2,3]
 * 输出: 6
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: [1,2,3,4]
 * 输出: 24
 * 
 * 
 * 注意:
 * 
 * 
 * 给定的整型数组长度范围是[3,10^4]，数组中所有的元素范围是[-1000, 1000]。
 * 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
     // 排序法
    // 执行用时：124 ms, 在所有 PHP 提交中击败了61.76%的用户
    // 内存消耗：15.8 MB, 在所有 PHP 提交中击败了20.00%的用户
    function maximumProduct1($nums) {
        // 数组长度最少也是三位数
        $size = count($nums);
        if($size == 3) {
            return $nums[0]*$nums[1]*$nums[2];
        }
        sort($nums); // 排序
        return max($nums[$size-1]*$nums[$size-2]*$nums[$size-3] , $nums[$size-1] * $nums[0] * $nums[1]);
    }

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 线性遍历
    // 执行用时：96 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.6 MB, 在所有 PHP 提交中击败了75.00%的用户
    function maximumProduct($nums) {
        // 数组长度最少也是三位数
        $size = count($nums);
        if($size == 3) {
            return $nums[0]*$nums[1]*$nums[2];
        }
        $min1 = $min2 = 1000;
        $max1 = $max2 = $max3 = -1000;
        foreach($nums as $item) {
            if($item <= $min1) {
                $min2 = $min1;
                $min1 = $item; // min1 最小
            }elseif($item <= $min2){
                $min2 = $item;
            }
            if($item >= $max1){
                $max3 = $max2;
                $max2 = $max1;
                $max1 = $item; // max1 最大
            }elseif($item >= $max2){
                $max3 = $max2;
                $max2 = $item;
            }elseif($item >= $max3){
                $max3 = $item;
            }
        }
        return max($max1*$max2*$max3, $min1*$min2*$max1);
    }
}
// @lc code=end

// @tacks think=start
/*
题意 meaning 
    这个题，有点类似昨天的有序数组平方（[977] 有序数组的平方），今天这个题是数组中三个数的最大乘积（[628] 三个数的最大乘积）
    
    其中有个共同的地方。
        有序数组的平方，要求平方后依然有序，（因为可能有负数），需要用到排序功能，或者双指针方法。
        三个数的乘积最大，那么直接就是想到排序，而且排序后看最大的几个数，然后拿到乘积，（负数需要考虑）。

    而其实很多题目我们最直观想到的就是排序，但是排序时间复杂度不一定能保证在O(n)。
        但是再一看，双指针的方法也就是线性扫描一遍即可拿到有序数组的排序。
        今天这个三个数乘积最大，是不是也可以不用排序？？？
    且听我一一分析！！！

关键 key 

    “三个数”，“乘积最大”

想法 idea

【1】排序法

    最直观的思路，基本不用多说！！！

    当我们给一个数组排序，按照升序排列。有下面几种情况。
    数组排序 =》 时间复杂度(nlogn)
        1> 后三位全是非负数，直接拿后三位进行乘积，即最大的。
        2> 后两位是非负数，  还是直接拿后三位进行乘积，即最大的。
        3> 后一位是非负数，  拿第一个第二个负数乘以最后面的，即最大的。
        4> 全部都是负数，    还是直接拿后三位进行乘积，即最大的。
    几种情况，再汇总一下，只有两种情况。
        ① 1、2、4> 直接拿后三位进行乘积
        ② 3>       最小的两个数，和最大的数进行乘积。
    =》 也就是比较①与②的大小。从而拿到最大值。

综上所述：该思路的 时间复杂度(nlogn)。

【2】线性扫描

    也是看到了题解，就是转变思路，直接拿到最大的三个数和最小的两个数，也就是定义五个变量，但是不用排序。

    然后再比较乘积最大的。

    min1 最小
    max1 最大
    min1,min2,max1,max2,max3 那么我们直接比较

    max1*max2*max3 与 min1*min2*max 即可拿到最大乘积。





*/
// @tacks think=end