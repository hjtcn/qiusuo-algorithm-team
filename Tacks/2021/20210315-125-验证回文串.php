<?php
/*
 * @Descripttion: 周一 今天北京沙尘暴很大
 * @Author: tacks321@qq.com
 * @Date: 2021-03-15 11:09:23
 * @LastEditTime: 2021-03-15 11:34:08
 */


/*
 * @lc app=leetcode.cn id=125 lang=php
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.81%)
 * Likes:    347
 * Dislikes: 0
 * Total Accepted:    208.9K
 * Total Submissions: 444K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 * 
 * 示例 1:
 * 
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "race a car"
 * 输出: false
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome($s) {
        $len  = strlen($s);
        $left = 0;
        $right= $len - 1;
        while($left < $right) {
            // 如果当前字符 不是数字或者字母 ，就跳过
            if(!ctype_alnum($s[$left])) {
                $left++;
                continue; // 重新判断
            }
            if(!ctype_alnum($s[$right])) {
                $right--;
                continue;
            }

            // 两个字符不相等
            if(strcasecmp($s[$left], $s[$right]) != 0) {
                return false; // 直接返回不是回文字符
            }
            // 继续向下比较
            $left++;
            $right--;
        }
        return true;
    }
}
// @lc code=end


/*
【题目】
    验证回文串
        只考虑字母和数字字符，可以忽略字母的大小写

        也就是说，不是字母或者数字的字符都直接跳过进行比较
【解析】

    回文 
        双指针比较法
        
        ctype_alnum ( string $text ) : bool
        如果text中所有的字符全部是字母和(或者)数字，返回 true 否则返回false

        strcasecmp — 二进制安全比较字符串（不区分大小写）。

    时间复杂度
        O(n)
*/