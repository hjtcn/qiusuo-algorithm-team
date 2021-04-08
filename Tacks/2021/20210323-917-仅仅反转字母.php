<?php
/*
 * @Descripttion: 周一 补一下题 最近很没意思哎
 * @Author: tacks321@qq.com
 * @Date: 2021-03-23 18:11:58
 * @LastEditTime: 2021-03-23 18:48:19
 */


/*
 * @lc app=leetcode.cn id=917 lang=php
 *
 * [917] 仅仅反转字母
 *
 * https://leetcode-cn.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (56.46%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    24.5K
 * Total Submissions: 43.3K
 * Testcase Example:  '"ab-cd"'
 *
 * 给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："ab-cd"
 * 输出："dc-ba"
 * 
 * 
 * 示例 2：
 * 
 * 输入："a-bC-dEf-ghIj"
 * 输出："j-Ih-gfE-dCba"
 * 
 * 
 * 示例 3：
 * 
 * 输入："Test1ng-Leet=code-Q!"
 * 输出："Qedo1ct-eeLg=ntse-T!"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122 
 * S 中不包含 \ or "
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $S
     * @return String
     */
    // 字母栈
    // 执行用时：4 ms, 在所有 PHP 提交中击败了90.00%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了70.00%的用户
    function reverseOnlyLetters($S) {
        $stack = [];
        $res = '';

        $len = strlen($S);
        // 入栈，记录所有的字母
        for ($i=0; $i < $len; $i++) { 
            if($this->_isLetter($S[$i])) {
                array_push($stack, $S[$i]);
            }
        }

        for ($i=0; $i < $len; $i++) { 
            // 如果是字母，出栈 （先进后出，所以就可以反转字母了）
            if($this->_isLetter($S[$i])) {
                $res .= array_pop($stack);
            } else {
            // 否则直接追加字符
                $res .= $S[$i];
            }
        }

        return $res;
        
    }

    // 判断是否是字母
    private function _isLetter($char) {
        if ( ($char >= 'a' && $char <= 'z' ) || ($char >= 'A' && $char <='Z' ) ) {
            return true;
        } else {
            return false;
        }
    }
}
// @lc code=end

/*
【题目】
    思考，首先要判断一个字符是否是字母

    然后单纯只反转字母，这个时候可以想到双指针解决，也可以想到栈先进后出
【思考】
    字母栈  
        将 s 中的所有字母单独存入栈中，所以出栈等价于对字母反序操作。
        然后，遍历 s 的所有字符，如果是字母我们就选择栈顶元素输出。
    时间复杂度
        时间复杂度：O(N)，其中 N 是 S 的长度。
        空间复杂度：O(N)。


 

*/