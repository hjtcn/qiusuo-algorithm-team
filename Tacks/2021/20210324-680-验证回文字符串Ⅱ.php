<?php
/*
 * @Descripttion: 周三了，马上一周就结束了
 * @Author: tacks321@qq.com
 * @Date: 2021-03-24 14:56:30
 * @LastEditTime: 2021-03-24 15:17:45
 */

/*
 * @lc app=leetcode.cn id=680 lang=php
 *
 * [680] 验证回文字符串 Ⅱ
 *
 * https://leetcode-cn.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (39.90%)
 * Likes:    336
 * Dislikes: 0
 * Total Accepted:    67.1K
 * Total Submissions: 167.4K
 * Testcase Example:  '"aba"'
 *
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 * 
 * 示例 1:
 * 
 * 
 * 输入: "aba"
 * 输出: True
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: "abca"
 * 输出: True
 * 解释: 你可以删除c字符。
 * 
 * 
 * 注意:
 * 
 * 
 * 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    // 双指针
    // 执行用时：52 ms, 在所有 PHP 提交中击败了70.37%的用户
    // 内存消耗：15.5 MB, 在所有 PHP 提交中击败了74.07%的用户
    function validPalindrome($s) {
        $len = strlen($s);
        // 双指针
        for ($i=0, $j=$len-1; $i < $len ; $i++,$j--) { 
            // 向里面收缩
            if($s[$i] != $s[$j]) {
                // 如果不相等
                // 有可能是左边字符错误，也有可能是右边字符错误
                return $this->_compare($s, $i+1, $j) || $this->_compare($s, $i, $j-1); 
            }
        }
        return true;
    }

    // 回文字符串验证
    private function _compare($s, $i, $j) {
        // 两边比较
        for(;$i<$j; $i++,$j--) {
            if($s[$i] != $s[$j]) {
                return false;
            }
        }
        return true;
    }
}
// @lc code=end

/*
【题目】
    非空字符串 s，最多删除一个字符。是否能成为回文字符串
    字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
        =>
            也就是说，可以不删除字符，或者删除一个字符
            字符串中只有 a b c ~ z

【题解】
    这个题 看似简单，但这个删除字符，看起来好难搞

    其实基本思路还是模拟回文字符串的比较过程
        就是利用双指针，然后不断向里面收缩

        因为可以去除一个字符，所以在不相等的时候，又一次试错的机会
            也就是双指针的一边进行移动，然后再次回文比较
*/