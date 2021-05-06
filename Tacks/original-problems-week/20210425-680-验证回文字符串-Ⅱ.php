<?php
/*
 * @Descripttion: 回文串
 * @Author: tacks321@qq.com
 * @Date: 2021-04-25 18:50:03
 * @LastEditTime: 2021-04-25 18:59:55
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
    // 执行用时：32 ms, 在所有 PHP 提交中击败了89.29%的用户
    // 内存消耗：15.5 MB, 在所有 PHP 提交中击败了78.57%用户
    function validPalindrome($s) {
        $len = strlen($s);
        $i = 0;
        $j = $len-1;
        while($i<$j) {
            // 向里收缩
            if($s[$i] != $s[$j]) {
                // 那么就可以两次试错
                return $this->_helper($s, $i+1, $j)
                || $this->_helper($s, $i, $j-1);
            }
            $i++;
            $j--;
        }
        return true;
    }

    // 子函数
    private function _helper($s, $i, $j) {
        while($i<=$j) {
            if($s[$i] != $s[$j]) {
                return false;
            }
            $i++;
            $j--;
        }
        return true;
    }
}
// @lc code=end

/*
【题目】
    给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

        就是有一次试错的机会，相当于我们玩的癞子
【题解】
    双指针搞
        模拟回文字符串比较过程
            双指针不断向里面收缩
       
    时间复杂度O(n)
    空间复杂度O(1)
        

*/