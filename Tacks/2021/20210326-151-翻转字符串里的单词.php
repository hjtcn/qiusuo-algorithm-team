<?php
/*
 * @Descripttion: 周五了，一周又要结束了，开森
 * @Author: tacks321@qq.com
 * @Date: 2021-03-26 11:46:05
 * @LastEditTime: 2021-03-26 19:22:30
 */


/*
 * @lc app=leetcode.cn id=151 lang=php
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (44.85%)
 * Likes:    300
 * Dislikes: 0
 * Total Accepted:    124.9K
 * Total Submissions: 271.1K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 * 
 * 说明：
 * 
 * 
 * 无空格字符构成一个 单词 。
 * 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："the sky is blue"
 * 输出："blue is sky the"
 * 
 * 
 * 示例 2：
 * 
 * 输入："  hello world!  "
 * 输出："world! hello"
 * 解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 
 * 
 * 示例 3：
 * 
 * 输入："a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 示例 4：
 * 
 * 输入：s = "  Bob    Loves  Alice   "
 * 输出："Alice Loves Bob"
 * 
 * 
 * 示例 5：
 * 
 * 输入：s = "Alice does not even like bob"
 * 输出："bob like even not does Alice"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 请尝试使用 O(1) 额外空间复杂度的原地解法。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $s
     * @return String
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了96.23%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了84.91%的用户
    function reverseWords($s) {
        $s  = trim($s);
        $res= '';
        $len= strlen($s);

        for ($i=$len-1; $i >0 ; $i--) { 
            if($s[$i] === ' ') {
                $res .= substr($s, $i+1, $len-$i-1). ' ';
                while($s[$i-1] === ' ') {
                    $i--;
                }
                $len = $i;
            }            
        }
        $res.= substr($s, 0, $len);
        return $res;
    }
}
// @lc code=end

/*
【题目】
    一个有空格的字符串，逐个翻转字符串中的每个单词。
        字符串 =>  包含英文大小写字母、数字和空格
        单词   =>  无空格字符构成一个单词 
        输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括 => 前后空格全部删除
        如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一         => 每个单词之间的空格只保留一个
        O(1) 额外空间复杂度的原地解法 => 不能使用哈希数组这种

    

*/