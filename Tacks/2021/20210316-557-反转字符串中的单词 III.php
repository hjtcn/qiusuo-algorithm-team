<?php

/*
 * @Descripttion: 周二了，今天空气还好一点，沙尘没有那么大
 * @Author: tacks321@qq.com
 * @Date: 2021-03-16 09:55:41
 * @LastEditTime: 2021-03-16 10:38:19
 */

/*
 * @lc app=leetcode.cn id=557 lang=php
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (73.71%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    113.8K
 * Total Submissions: 154.1K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入："Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return String
     */
    // 【暴力解决】
    // 执行用时：28 ms, 在所有 PHP 提交中击败了53.13%的用户
    // 内存消耗：15.6 MB, 在所有 PHP 提交中击败了37.50%的用户
    function reverseWords($s) {
        $res = ''; // 结果字符串
        
        // 字符串按照空格分隔数组
        $arr = explode(' ', $s);
        foreach($arr as $item) {
            $res .= $this->_reverseItem($item) . ' ';
        }
        // 最后把空格去掉
        $res = trim($res);
        return $res;
    }

    /**
     * 对每一个字符串进行反转
     * @param [string] $item
     */
    function _reverseItem($item) {
        $str = '';
        $len = strlen($item);
        // 遍历字符串 反向拼接
        for ($i=$len-1; $i >= 0; $i--) { 
            $str = $str.$item[$i];
        }

        return $str;
    }
}


class Solution2 {

    /**
     * @param String $s
     * @return String
     */
    // 【就地解决】
    // 执行用时：44 ms, 在所有 PHP 提交中击败了21.88%的用户
    // 内存消耗：15.6 MB, 在所有 PHP 提交中击败了62.50%的用户
    function reverseWords($s) {
        $len = strlen($s);
        $i   = 0;

        // 从头遍历字符串
        while($i < $len) {
            $start = $i; // 记录第一个子串的第一个位置
            //  先后遍历到第一个子串 结束的时候
            while($i < $len && $s[$i] != ' ') {
                $i++;
            }

            // 子串 起始位置
            $left = $start;
            $right= $i - 1;

            // 字串的反转
            while($left < $right) {
                // 交换两个变量
                list($s[$right], $s[$left]) = array($s[$left], $s[$right]); // list() 函数用数组中的元素为一组变量赋值
                $left++;
                $right--;
            }

            // 等于空格的时候，继续向后遍历
            while($i < $len && $s[$i] == ' ') {
                $i++;
            }

        }   
        return $s;
    }

}


class Solution3 {

    /**
     * @param String $s
     * @return String
     */
    // 【函数法】
    // 执行用时：8 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了87.50%的用户
    function reverseWords($s) {
        // 字符串分割为数组
        $arr = explode(' ', $s);
        // 直接改变数字的元素
        foreach($arr as &$value) {
            $value = strrev($value); // 字符串反转
        }
        unset($value); // 释放变量
        // 数组元素合并成字符串
        $s = implode(' ', $arr);
        return $s;
    }

}



$obj = new Solution3;

var_dump($obj->reverseWords("Let's take LeetCode contest")) ;


// @lc code=end

/*
【题目】
    你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序

    也就是一句话中，对每个英语单词进行反转，然后保留对应的空格
【解题】

1、暴力解决
    我一开始的想法，用 explode 函数（ 使用 空格 分割另一个字符串，然后成为一个字符串数组）
    然后对字符串数组进行循环处理每一个item进行 反转。

    时间复杂度就是遍历整个字符串 O(n)
    空间复杂度申请了数组 O(n) 基本上就是(N/空格个数)

2、就地解决
    优点: 避免开辟新的空间
    缺点: 数据类型的限制，字符串是否可以变化

    时间复杂度O(n) (像这种必须要遍历所有的字符的，肯定时间复杂度最低也是O(n)了吧)
    空间复杂度O(1)

3、函数法

    众所周知，PHP的函数是很牛逼的。

    所以实际上如果业务中需要这样做，肯定选择第三种方案，虽然申请额外空间，但是函数处理速度会比你单纯for好一些 （貌似


    
*/