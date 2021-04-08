<?php
/*
 * @Descripttion: 周四的题，提前干掉
 * @Author: tacks321@qq.com
 */



/*
 * @lc app=leetcode.cn id=647 lang=php
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (64.71%)
 * Likes:    522
 * Dislikes: 0
 * Total Accepted:    86.1K
 * Total Submissions: 132.3K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
 * 
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 * 
 * 
 * 示例 2：
 * 
 * 输入："aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的字符串长度不会超过 1000 。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    // 中心扩展法
    // 执行用时：40 ms, 在所有 PHP 提交中击败了88.46%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了100.00%的用户
    function countSubstrings($s) {
        $num = 0;
        $len = strlen($s);

        // 遍历回文中心点
        for ($i=0; $i < $len; $i++) { 
            // j=0 中心是一个点  j=1 中心是两个点,其实就是奇数or偶数
            for ($j=0; $j <= 1; $j++) { 
                $left = $i;
                $right= $i+$j;
                while($left>=0 && $right<$len && $s[$left] == $s[$right]) {
                    // 向两边扩展
                    $left--;
                    $right++;
                    $num++;

                }
            }
        }
        return $num;

    }
}
// @lc code=end

/*
【题目】
    字符串，计算这个字符串中有多少个回文子串
    => 又是字符串，判断回文，这个回文出现次数挺高阿

    具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串
    => 也就是说，这次回文的唯一性判断，是根据位置识别的
【解析】
    中心拓展法

        看题解实现的，还是有些细节的地方不太明白

    貌似动态规划什么的也能做


*/